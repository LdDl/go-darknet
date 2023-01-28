package darknet

// #cgo !windows pkg-config: opencv4
// #cgo CXXFLAGS:   --std=c++1z
// #cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
// #cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core420 -lopencv_face420 -lopencv_videoio420 -lopencv_imgproc420 -lopencv_highgui420 -lopencv_imgcodecs420 -lopencv_objdetect420 -lopencv_features2d420 -lopencv_video420 -lopencv_dnn420 -lopencv_xfeatures2d420 -lopencv_plot420 -lopencv_tracking420 -lopencv_img_hash420 -lopencv_calib3d420
// #include <stdlib.h>
// #include <stdint.h>
// #include <stdlib.h>
//
// #include <darknet.h>
//
// #include "yolo.h"
import "C"
import (
	"image"
	"unsafe"
)

type BBOX struct {
	Rect          image.Rectangle // (x,y) - top-left corner, (w, h) - width & height of bounded box
	Confidence    float32         // confidence - probability that the object was found correctly
	ObjId         uint            // class of object - from range [0, classes-1]
	TrackId       uint            // tracking id for video (0 - untracked, 1 - inf - tracked object)
	FramesCounter uint            // counter of frames on which the object was detected
	X3d, Y3d, Z3d float32         // center of object (in Meters) if ZED 3D Camera is used
}

func Init(cfg, weights string, gpu int) {
	ccfg := C.CString(cfg)
	cweights := C.CString(weights)
	defer C.free(unsafe.Pointer(ccfg))
	defer C.free(unsafe.Pointer(cweights))

	C.init(ccfg, cweights, C.int(gpu))
}

func GetDeviceCount() int {
	return int(C.get_device_count())
}

func GetDeviceName(gpu int, name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.get_device_name(C.int(gpu), cname))
}

func Dispose() int {
	return int(C.dispose())
}

func Detect(name string) (result int, bboxs []BBOX) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	bbox_t_container := C.bbbox_t_container{}

	result = int(C.detect_image(cname, &bbox_t_container))

	for i := 0; i < result; i++ {
		candidate := bbox_t_container.candidates[i]
		bbox := BBOX{
			image.Rect(int(candidate.x), int(candidate.y), int(candidate.w), int(candidate.h)),
			float32(candidate.prob),
			uint(candidate.obj_id),
			uint(candidate.track_id),
			uint(candidate.frames_counter),
			float32(candidate.x_3d),
			float32(candidate.y_3d),
			float32(candidate.z_3d),
		}

		bboxs = append(bboxs, bbox)
	}

	return
}

func DetectImageMat(data []byte) (result int, bboxs []BBOX) {
	bbox_t_container := C.bbbox_t_container{}
	result = int(C.detect_image_mat((*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)), &bbox_t_container))

	for i := 0; i < result; i++ {
		candidate := bbox_t_container.candidates[i]
		bbox := BBOX{
			image.Rect(int(candidate.x), int(candidate.y), int(candidate.x+candidate.w), int(candidate.y+candidate.h)),
			float32(candidate.prob),
			uint(candidate.obj_id),
			uint(candidate.track_id),
			uint(candidate.frames_counter),
			float32(candidate.x_3d),
			float32(candidate.y_3d),
			float32(candidate.z_3d),
		}

		bboxs = append(bboxs, bbox)
	}

	return
}
