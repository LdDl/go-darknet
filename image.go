package darknet

// #include <darknet.h>
import (
	"C"
	"image"
	"unsafe"
)

// DarknetImage represents the image buffer.
// type DarknetImage struct {
// 	Width  int
// 	Height int
// 	image  C.image
// }

func float_p(arr []float32) *C.float {
	return (*C.float)(unsafe.Pointer(&arr[0]))
}

// Image2Float32 Returns []float32 representation of image.Image
func Image2Float32(img image.Image) ([]float32, error) {
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

	return ans, nil
}
