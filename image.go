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

func float_p(arr []float32) *C.float {
	return (*C.float)(unsafe.Pointer(&arr[0]))
}

// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image/59747737#59747737
func image_2_array_pix(src image.Image) []float32 {
	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	src_rgba := image.NewRGBA(src.Bounds())
	draw.Copy(src_rgba, image.Point{}, src, src.Bounds(), draw.Src, nil)
	ans := []float32{}
	red := []float32{}
	green := []float32{}
	blue := []float32{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			rpix, gpix, bpix := float32(pix[0])/257.0, float32(pix[1])/257.0, float32(pix[2])/257.0
			red = append(red, rpix)
			green = append(green, gpix)
			blue = append(blue, bpix)
		}
	}
	ans = append(ans, red...)
	ans = append(ans, green...)
	ans = append(ans, blue...)
	return ans
}

// Image2Float32 Returns []float32 representation of image.Image
func Image2Float32(img image.Image) (*DarknetImage, error) {
	ans := image_2_array_pix(img)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	imgDarknet := &DarknetImage{
		Width:  width,
		Height: height,
		image:  C.make_image(C.int(width), C.int(height), 3),
	}
	// for i := range ans {
	// 	C.set_data_f32_val(imgDarknet.image.data, C.int(i), C.float(ans[i]))
	// }
	C.fill_image_f32(&imgDarknet.image, C.int(width), C.int(height), 3, float_p(ans))
	// imgDarknet.image = C.load_image_color(C.CString("/home/dimitrii/Downloads/mega.jpg"), 4032, 3024)
	return imgDarknet, nil
}