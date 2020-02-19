package darknet

// #include <stdlib.h>
// #include "image.h"
import "C"
import (
	"image"
	"unsafe"
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

func float_p(arr []float32) *C.float {
	return (*C.float)(unsafe.Pointer(&arr[0]))
}

// Image2Float32 Returns []float32 representation of image.Image
func Image2Float32(img image.Image) (DarknetImage, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	imgwh := width * height
	imgSize := imgwh * 3

	ans := make([]float32, imgSize)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(y, x).RGBA()
			rpix, gpix, bpix := float32(r>>8)/float32(255.0), float32(g>>8)/float32(255.0), float32(b>>8)/float32(255.0)
			ans[y+x*height] = rpix
			ans[y+x*height+imgwh] = gpix
			ans[y+x*height+imgwh+imgwh] = bpix
		}
	}

	imgDarknet := DarknetImage{
		Width:  width,
		Height: height,
		// image:  C.new_darknet_image(),
	}

	// imgDarknet.image = C.prepare_image(imgDarknet.image, C.int(width), C.int(height), 3)
	// imgDarknet.image.data = float_p(ans)
	// imgDarknet.image = C.resize_image(imgDarknet.image, 416, 416) // Do we need resize? (detection function actually does it)

	imgDarknet.image = C.load_image_color(C.CString("/home/dimitrii/Downloads/mega.jpg"), 416, 416)

	return imgDarknet, nil
}
