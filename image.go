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
	ans    []float32
	image  C.image
}

// Close and release resources.
func (img *DarknetImage) Close() error {
	C.free_image(img.image)
	img.ans = nil
	return nil
}

// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image/59747737#59747737
func imgTofloat32(src image.Image) []float32 {
	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	srcRGBA := image.NewRGBA(src.Bounds())
	draw.Copy(srcRGBA, image.Point{}, src, src.Bounds(), draw.Src, nil)

	red := make([]float32, 0, width*height)
	green := make([]float32, 0, width*height)
	blue := make([]float32, 0, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idxSource := (y*width + x) * 4
			pix := srcRGBA.Pix[idxSource : idxSource+4]
			rpix, gpix, bpix := float32(pix[0])/257.0, float32(pix[1])/257.0, float32(pix[2])/257.0
			red = append(red, rpix)
			green = append(green, gpix)
			blue = append(blue, bpix)
		}
	}
	srcRGBA = nil

	ans := make([]float32, len(red)+len(green)+len(blue))
	copy(ans[:len(red)], red)
	copy(ans[len(red):len(red)+len(green)], green)
	copy(ans[len(red)+len(green):], blue)
	red = nil
	green = nil
	blue = nil
	return ans
}

// Image2Float32 Returns []float32 representation of image.Image
func Image2Float32(img image.Image) (*DarknetImage, error) {
	// ans := imgTofloat32(img)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	imgDarknet := &DarknetImage{
		Width:  width,
		Height: height,
		ans:    imgTofloat32(img),
		image:  C.make_image(C.int(width), C.int(height), 3),
	}
	C.fill_image_f32(&imgDarknet.image, C.int(width), C.int(height), 3, (*C.float)(unsafe.Pointer(&imgDarknet.ans[0])))
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
