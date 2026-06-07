package capture

import (
	"image"

	"github.com/go-vgo/robotgo"
)

func CaptureScreen(x, y, w, h int) image.Image {
	bit := robotgo.CaptureScreen(x, y, w, h)
	defer robotgo.FreeBitmap(bit)
	img := robotgo.ToImage(bit)
	return img
}
