package darknet

// #include <darknet.h>
import "C"
import "errors"

// Image represents the image buffer.
type Image struct {
	Width  int
	Height int

	image C.image
}

var errUnableToLoadImage = errors.New("unable to load image")

// Close and release resources.
func (img *Image) Close() error {
	C.free_image(img.image)
	return nil
}

// ImageFromPath reads image file specified by path.
func ImageFromPath(path string) (*Image, error) {
	img := Image{
		image: C.load_image_color(C._GoStringPtr(path), 0, 0),
	}

	if img.image.data == nil {
		return nil, errUnableToLoadImage
	}

	img.Width = int(img.image.w)
	img.Height = int(img.image.h)
	return &img, nil
}
