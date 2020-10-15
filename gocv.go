// +build gocv

package darknet

// #cgo !windows pkg-config: opencv4
// #cgo CXXFLAGS:   --std=c++1z
// #cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
// #cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_core420 -lopencv_face420 -lopencv_videoio420 -lopencv_imgproc420 -lopencv_highgui420 -lopencv_imgcodecs420 -lopencv_objdetect420 -lopencv_features2d420 -lopencv_video420 -lopencv_dnn420 -lopencv_xfeatures2d420 -lopencv_plot420 -lopencv_tracking420 -lopencv_img_hash420 -lopencv_calib3d420
// #include <stdlib.h>
// #include <stdint.h>
// #include "yolo.h"
import "C"
import (
	"image"
	"unsafe"

	"gocv.io/x/gocv"
)

func DetectMat(mat gocv.Mat) (result int, bboxs []BBOX) {
	bbox_t_container := C.bbbox_t_container{}
	result = int(C.detect_mat(unsafe.Pointer(mat.Ptr()), &bbox_t_container))

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
