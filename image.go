package darknet

// #include <darknet.h>
import "C"
import (
	"errors"
	"unsafe"
)

// Image represents the image buffer.
type Image struct {
	Width  int
	Height int

	image C.image
}

var (
	errUnableToLoadImage = errors.New("unable to load image")
)

// Close and release resources.
func (img *Image) Close() error {
	C.free_image(img.image)
	return nil
}

// ImageFromPath reads image file specified by path.
func ImageFromPath(path string) (*Image, error) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	img := Image{
		image: C.load_image_color(p, 0, 0),
	}

	if img.image.data == nil {
		return nil, errUnableToLoadImage
	}

	img.Width = int(img.image.w)
	img.Height = int(img.image.h)

	return &img, nil
}

// ImageFromMemory reads image file data represented by the specified byte
// slice.
func ImageFromMemory(buf []byte) (*Image, error) {
	cBuf := C.CBytes(buf)
	defer C.free(cBuf)

	img := Image{
		image: C.load_image_from_memory_color((*C.uchar)(cBuf),
			C.int(len(buf)), 0, 0),
	}

	if img.image.data == nil {
		return nil, errUnableToLoadImage
	}

	img.Width = int(img.image.w)
	img.Height = int(img.image.h)

	return &img, nil
}
