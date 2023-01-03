package darknet

// #include <stdlib.h>
//
// #include <darknet.h>
//
// #include "network.h"
import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

// YOLONetwork represents a neural network using YOLO.
type YOLONetwork struct {
	GPUDeviceIndex           int
	NetworkConfigurationFile string
	WeightsFile              string
	Threshold                float32

	ClassNames []string
	Classes    int

	cNet                *C.network
	hierarchalThreshold float32
	nms                 float32
}

var (
	errNetworkNotInit      = errors.New("Network not initialised")
	errUnableToInitNetwork = errors.New("Unable to initialise")
)

// Init the network.
func (n *YOLONetwork) Init() error {
	nCfg := C.CString(n.NetworkConfigurationFile)
	defer C.free(unsafe.Pointer(nCfg))
	wFile := C.CString(n.WeightsFile)
	defer C.free(unsafe.Pointer(wFile))
	// GPU device ID must be set before `load_network()` is invoked.
	C.cuda_set_device(C.int(n.GPUDeviceIndex))
	n.cNet = C.load_network(nCfg, wFile, 0)
	if n.cNet == nil {
		return errUnableToInitNetwork
	}
	C.srand(2222222)
	n.hierarchalThreshold = 0.5
	n.nms = 0.45
	metadata := C.get_metadata(nCfg)
	n.Classes = int(metadata.classes)
	n.ClassNames = makeClassNames(metadata.names, n.Classes)
	return nil
}

/* EXPERIMENTAL */
/*
	By default AlexeyAB's Darknet doesn't export any functions in darknet.h to give ability to create network from scratch via code.
	So I can't modify `parse_network_cfg_custom` to load `list *sections = read_cfg(filename);` from memory.

	So, the point of this method is to be able create network configuration via Golang and then pass it to `C.load_network`
*/
func (n *YOLONetwork) InitFromDefinedCfg() error {
	wFile := C.CString(n.WeightsFile)
	defer C.free(unsafe.Pointer(wFile))
	/* Prepare network sections via Go */
	/*
		instead of using:
			wFile := C.CString(n.WeightsFile)
			defer C.free(unsafe.Pointer(wFile))
		We call `load_network` that takes the first char* parameter (representing a file path to network configuration) with a Go function that takes an in-memory file
	*/
	cfgBytes, err := os.ReadFile(n.NetworkConfigurationFile)
	if err != nil {
		return errors.Wrap(err, "Can't read file bytes")
	}
	// Create a temporary file.
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		return errors.Wrap(err, "Can't create temporary file")
	}
	defer os.Remove(tmpFile.Name())
	fmt.Println("here", tmpFile.Name())
	// Write the file content to the temporary file.
	if _, err := tmpFile.Write(cfgBytes); err != nil {
		return errors.Wrap(err, "Can't write network's configuration into temporary file")
	}
	defer tmpFile.Close()
	// Open the temporary file.
	fd, err := syscall.Open(tmpFile.Name(), syscall.O_RDWR, 0)
	if err != nil {
		return errors.Wrap(err, "Can't re-open temporary file")
	}
	defer syscall.Close(fd)
	// Create a memory mapping of the file.
	addr, err := syscall.Mmap(fd, 0, len(cfgBytes), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return errors.Wrap(err, "Can't mmap on temporary file")
	}

	// GPU device ID must be set before `load_network()` is invoked.
	C.cuda_set_device(C.int(n.GPUDeviceIndex))
	nCfg := C.CString(tmpFile.Name())
	defer C.free(unsafe.Pointer(nCfg))
	n.cNet = C.load_network(nCfg, wFile, 0)
	if n.cNet == nil {
		return errUnableToInitNetwork
	}
	C.srand(2222222)
	n.hierarchalThreshold = 0.5
	n.nms = 0.45
	metadata := C.get_metadata(nCfg)
	n.Classes = int(metadata.classes)
	n.ClassNames = makeClassNames(metadata.names, n.Classes)
	// Unmap the memory-mapped file.
	if err := syscall.Munmap(addr); err != nil {
		return errors.Wrap(err, "Can't revert mmap")
	}
	return nil
}

// Close and release resources.
func (n *YOLONetwork) Close() error {
	if n.cNet == nil {
		return errNetworkNotInit
	}
	C.free_network_ptr(n.cNet)
	n.cNet = nil
	return nil
}

// Detect specified image
func (n *YOLONetwork) Detect(img *DarknetImage) (*DetectionResult, error) {
	if n.cNet == nil {
		return nil, errNetworkNotInit
	}
	startTime := time.Now()
	result := C.perform_network_detect(n.cNet, &img.image, C.int(n.Classes), C.float(n.Threshold), C.float(n.hierarchalThreshold), C.float(n.nms), C.int(0))
	endTime := time.Now()
	ds := makeDetections(img, result.detections, int(result.detections_len), n.Threshold, n.Classes, n.ClassNames)
	C.free_detections(result.detections, result.detections_len)
	endTimeOverall := time.Now()
	out := DetectionResult{
		Detections:           ds,
		NetworkOnlyTimeTaken: endTime.Sub(startTime),
		OverallTimeTaken:     endTimeOverall.Sub(startTime),
	}
	return &out, nil
}
