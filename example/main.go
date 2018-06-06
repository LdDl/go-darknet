package main

import (
	"flag"
	"fmt"
	"log"

	darknet "github.com/gyonluks/go-darknet"
)

var dataConfigFile = flag.String("dataConfigFile", "",
	"Path to data configuration file. Example: cfg/coco.data")
var configFile = flag.String("configFile", "",
	"Path to network layer configuration file. Example: cfg/yolov3.cfg")
var weightsFile = flag.String("weightsFile", "",
	"Path to weights file. Example: yolov3.weights")
var imageFile = flag.String("imageFile", "",
	"Path to image file, for detection. Example: image.jpg")

func printError(err error) {
	log.Println("error:", err)
}

func main() {
	flag.Parse()

	if *dataConfigFile == "" || *configFile == "" || *weightsFile == "" ||
		*imageFile == "" {
		flag.Usage()
		return
	}

	n := darknet.YOLONetwork{
		DataConfigurationFile:    *dataConfigFile,
		NetworkConfigurationFile: *configFile,
		WeightsFile:              *weightsFile,
		Threshold:                .5,
	}
	if err := n.Init(); err != nil {
		printError(err)
		return
	}
	defer n.Close()

	img, err := darknet.ImageFromPath(*imageFile)
	if err != nil {
		printError(err)
		return
	}
	defer img.Close()

	dr, err := n.Detect(img)
	if err != nil {
		printError(err)
		return
	}

	log.Println("Network-only time taken:", dr.NetworkOnlyTimeTaken)
	log.Println("Overall time taken:", dr.OverallTimeTaken)
	for _, d := range dr.Detections {
		for i := range d.ClassIDs {
			bBox := d.BoundingBox
			fmt.Printf("%s (%d): %.4f%% | start point: (%f,%f) | height: %f | width: %f\n",
				d.ClassNames[i], d.ClassIDs[i],
				d.Probabilities[i],
				bBox.X, bBox.Y, bBox.Height, bBox.Width,
			)
		}
	}
}
