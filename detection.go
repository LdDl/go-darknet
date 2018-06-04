package darknet

// #include <darknet.h>
//
// #include "detection.h"
import "C"

// Detection represents a detection.
type Detection struct {
	BoundingBox

	ClassIDs      []int
	ClassNames    []string
	Probabilities []float32
}

func makeDetection(det *C.detection, threshold float32, classes int,
	classNames []string) *Detection {
	dClassIDs := make([]int, 0)
	dClassNames := make([]string, 0)
	dProbs := make([]float32, 0)

	for i := 0; i < int(classes); i++ {
		dProb := float32(
			C.get_detection_probability(det, C.int(i), C.int(classes)))
		if dProb > threshold {
			dClassIDs = append(dClassIDs, i)
			cN := classNames[i]
			dClassNames = append(dClassNames, cN)
			dProbs = append(dProbs, dProb*100)
		}
	}

	out := Detection{
		BoundingBox: BoundingBox{
			X:      float32(det.bbox.x),
			Y:      float32(det.bbox.y),
			Width:  float32(det.bbox.w),
			Height: float32(det.bbox.h),
		},

		ClassIDs:      dClassIDs,
		ClassNames:    dClassNames,
		Probabilities: dProbs,
	}

	return &out
}

func makeDetections(detections *C.detection, detectionsLength int,
	threshold float32, classes int, classNames []string) []*Detection {
	// Make list of detection objects.
	ds := make([]*Detection, detectionsLength)
	for i := 0; i < int(detectionsLength); i++ {
		det := C.get_detection(detections, C.int(i), C.int(classes))
		d := makeDetection(det, threshold, classes, classNames)
		ds[i] = d
	}

	return ds
}
