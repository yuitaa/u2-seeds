package capture

import (
	"image"

	"github.com/go-vgo/robotgo"
)

func CaptureScreen(r Rect) image.Image {
	bit := robotgo.CaptureScreen(r.X, r.Y, r.W, r.H)
	defer robotgo.FreeBitmap(bit)
	img := robotgo.ToImage(bit)
	return img
}
