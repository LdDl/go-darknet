package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	darknet "github.com/LdDl/go-darknet"
)

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

	if *configFile == "" || *weightsFile == "" ||
		*imageFile == "" {

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
		printError(err)
		return
	}
	defer n.Close()

	infile, err := os.Open(*imageFile)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()
	src, err := jpeg.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	imgDarknet, err := darknet.Image2Float32(src)
	if err != nil {
		panic(err.Error())
	}
	defer imgDarknet.Close()

	dr, err := n.Detect(imgDarknet)
	if err != nil {
		printError(err)
		return
	}

	log.Println("Network-only time taken:", dr.NetworkOnlyTimeTaken)
	log.Println("Overall time taken:", dr.OverallTimeTaken, len(dr.Detections))
	for _, d := range dr.Detections {
		for i := range d.ClassIDs {
			bBox := d.BoundingBox
			fmt.Printf("%s (%d): %.4f%% | start point: (%d,%d) | end point: (%d, %d)\n",
				d.ClassNames[i], d.ClassIDs[i],
				d.Probabilities[i],
				bBox.StartPoint.X, bBox.StartPoint.Y,
				bBox.EndPoint.X, bBox.EndPoint.Y,
			)
		}
	}
}

func imageToBytes(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, nil)
	return buf.Bytes(), err
}
