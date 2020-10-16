package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LdDl/go-darknet"
)

var configFile = flag.String("configFile", "",
	"Path to network layer configuration file. Example: cfg/yolov3.cfg")
var weightsFile = flag.String("weightsFile", "",
	"Path to weights file. Example: yolov3.weights")
var serverPort = flag.Int("port", 8090,
	"Listening port")

func main() {
	flag.Parse()

	if *configFile == "" || *weightsFile == "" {
		flag.Usage()
		return
	}

	n := darknet.YOLONetwork{
		GPUDeviceIndex:           0,
		NetworkConfigurationFile: *configFile,
		WeightsFile:              *weightsFile,
		Threshold:                .25,
	}
	if err := n.Init(); err != nil {
		log.Println(err)
		return
	}
	defer n.Close()

	http.HandleFunc("/detect_objects", detectObjects(&n))
	http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil)
}

// DarknetResp Response
type DarknetResp struct {
	NetTime       string              `json:"net_time"`
	OverallTime   string              `json:"overall_time"`
	NumDetections int                 `json:"num_detections"`
	Detections    []*DarknetDetection `json:"detections"`
}

// DarknetDetection Information about single detection
type DarknetDetection struct {
	ClassID     int           `json:"class_id"`
	ClassName   string        `json:"class_name"`
	Probability float32       `json:"probability"`
	StartPoint  *DarknetPoint `json:"start_point"`
	EndPoint    *DarknetPoint `json:"end_point"`
}

// DarknetPoint image.Image point
type DarknetPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func detectObjects(n *darknet.YOLONetwork) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// Restrict file size up to 10mb
		req.ParseMultipartForm(10 << 20)

		file, _, err := req.FormFile("image")
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("Error reading FormFile: %s", err.Error()))
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("Error reading bytes: %s", err.Error()))
			return
		}

		imgSrc, _, err := image.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("Error decoding bytes to image: %s", err.Error()))
			return
		}

		imgDarknet, err := darknet.Image2Float32(imgSrc)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("Error converting image.Image to darknet.DarknetImage: %s", err.Error()))
			return
		}
		defer imgDarknet.Close()

		dr, err := n.Detect(imgDarknet)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("Error detecting objects: %s", err.Error()))
			return
		}

		resp := DarknetResp{
			NetTime:       fmt.Sprintf("%v", dr.NetworkOnlyTimeTaken),
			OverallTime:   fmt.Sprintf("%v", dr.OverallTimeTaken),
			NumDetections: len(dr.Detections),
			Detections:    []*DarknetDetection{},
		}

		for _, d := range dr.Detections {
			for i := range d.ClassIDs {
				bBox := d.BoundingBox
				resp.Detections = append(resp.Detections, &DarknetDetection{
					ClassID:     d.ClassIDs[i],
					ClassName:   d.ClassNames[i],
					Probability: d.Probabilities[i],
					StartPoint: &DarknetPoint{
						X: bBox.StartPoint.X,
						Y: bBox.StartPoint.Y,
					},
					EndPoint: &DarknetPoint{
						X: bBox.EndPoint.X,
						Y: bBox.EndPoint.Y,
					},
				})
			}
		}

		respBytes, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("Error encoding response: %s", err.Error()))
			return
		}

		fmt.Fprintf(w, string(respBytes))
	}
}
