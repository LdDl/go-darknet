package darknet

// #include <stdlib.h>
// #include "image.h"
import "C"
import (
	"image"
	"unsafe"

	"golang.org/x/image/draw"
)

// DarknetImage represents the image buffer.
type DarknetImage struct {
	Width  int
	Height int
	image  C.image
}

// Close and release resources.
func (img *DarknetImage) Close() error {
	C.free_image(img.image)
	return nil
}

func Image2Float32(src image.Image) (*DarknetImage, error) {
	// Fast path: NRGBA → darknet CHW float32 directly, no intermediate alloc
	if nrgba, ok := src.(*image.NRGBA); ok {
		return ImageNRGBA2Float32(nrgba)
	}
	// Fast path: RGBA
	if rgba, ok := src.(*image.RGBA); ok {
		return ImageRGBA2Float32(rgba)
	}
	// Fallback: generic image → RGBA via draw.Copy
	bounds := src.Bounds()
	srcRGBA := image.NewRGBA(bounds)
	draw.Copy(srcRGBA, image.Point{}, src, bounds, draw.Src, nil)
	return ImageRGBA2Float32(srcRGBA)
}

// ImageNRGBA2Float32 converts *image.NRGBA directly to darknet image
// without intermediate RGBA allocation.
func ImageNRGBA2Float32(img *image.NRGBA) (*DarknetImage, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	imgDarknet := &DarknetImage{
		Width:  width,
		Height: height,
		image:  C.make_image(C.int(width), C.int(height), 3),
	}
	C.nrgba_to_float_and_fill_image(&imgDarknet.image, C.int(width), C.int(height),
		(*C.uint8_t)(unsafe.Pointer(&img.Pix[0])), C.int(img.Stride))
	return imgDarknet, nil
}

// Image2Float32 Returns []float32 representation of image.Image
func ImageRGBA2Float32(img *image.RGBA) (*DarknetImage, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	imgDarknet := &DarknetImage{
		Width:  width,
		Height: height,
		image:  C.make_image(C.int(width), C.int(height), 3),
	}
	C.to_float_and_fill_image(&imgDarknet.image, C.int(width), C.int(height), (*C.uint8_t)(unsafe.Pointer(&img.Pix[0])))
	return imgDarknet, nil
}

// Float32ToDarknetImage Converts []float32 to darknet image
func Float32ToDarknetImage(flatten []float32, width, height int) (*DarknetImage, error) {
	imgDarknet := &DarknetImage{
		Width:  width,
		Height: height,
		image:  C.make_image(C.int(width), C.int(height), 3),
	}
	C.fill_image_f32(&imgDarknet.image, C.int(width), C.int(height), 3, (*C.float)(unsafe.Pointer(&flatten[0])))
	return imgDarknet, nil
}
