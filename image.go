package darknet

// #include <stdlib.h>
// #include "image.h"
import "C"
import (
	"fmt"
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
func Image2Float32(img image.Image) (*DarknetImage, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	ans := []float32{}
	red := []float32{}
	green := []float32{}
	blue := []float32{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(y, x).RGBA()
			rpix, gpix, bpix := float32(r>>8)/float32(255.0), float32(g>>8)/float32(255.0), float32(b>>8)/float32(255.0)
			red = append(red, rpix)
			green = append(green, gpix)
			blue = append(blue, bpix)
		}
	}
	ans = append(ans, red...)
	ans = append(ans, green...)
	ans = append(ans, blue...)

	imgDarknet := &DarknetImage{
		Width:  width,
		Height: height,
		image:  C.make_image(C.int(width), C.int(height), 3),
	}
	// for i := range ans {
	// 	C.set_data_f32_val(imgDarknet.image.data, C.int(i), C.float(ans[i]))
	// }
	fmt.Println(ans[:100])
	// C.fill_image_f32(&imgDarknet.image, C.int(width), C.int(height), 3, float_p(ans))
	// newImg := C.resize_image_golang(imgDarknet.image, 416, 416)
	// r, g, b, _ := img.At(0, 0).RGBA()
	// rpix, gpix, bpix := float32(r>>8)/float32(255.0), float32(g>>8)/float32(255.0), float32(b>>8)/float32(255.0)
	// fmt.Println(rpix, gpix, bpix)

	// imgDarknet.image = C.load_image_color(C.CString("/home/dimitrii/Downloads/mega.jpg"), 4032, 3024)
	return imgDarknet, nil
}

// ImageFromMemory reads image file data represented by the specified byte
// slice.
func ImageFromMemory(buf []byte, width, height int) (*DarknetImage, error) {
	cBuf := C.CBytes(buf)
	defer C.free(cBuf)

	imgd := C.make_image(C.int(width), C.int(height), 3)
	// var imgd C.image
	// imgd.w = width
	// imgd.h = height
	// imgd.c = 3

	C.copy_image_from_bytes(imgd, (*C.char)(cBuf))
	img := DarknetImage{
		image: imgd,
	}

	if img.image.data == nil {
		return nil, fmt.Errorf("nil image")
	}

	img.Width = int(img.image.w)
	img.Height = int(img.image.h)

	return &img, nil
}
