package capture

import (
	"fmt"
	"image"

	"github.com/go-vgo/robotgo"
	"github.com/yuitaa/u2-seeds/internal/matching"
	"github.com/yuitaa/u2-seeds/internal/templates"
)

func GetInformation() {
	getTasks := func(r rect) (string, string) {
		screen := captureScreen(r)
		rgba, ok := screen.(*image.RGBA)
		if !ok {
			return "", ""
		}
		taskMatch := *matching.MatchTemplate(rgba, &templates.TaskTemplate, true)
		extensionMatch := *matching.MatchTemplate(rgba, &templates.ExtensionTemplate, true)

		var task string
		if len(taskMatch) > 0 {
			task = taskMatch[0].Key
		}

		var extension string
		if len(extensionMatch) > 0 {
			extension = extensionMatch[0].Key
		}

		return task, extension
	}

	upperTask, upperExtension := getTasks(upperTaskRect)
	lowerTask, lowerExtension := getTasks(lowerTaskRect)

	fmt.Printf("Task: %s, Extension: %s\n", upperTask, upperExtension)
	fmt.Printf("Task: %s, Extension: %s\n", lowerTask, lowerExtension)

	screen := captureScreen(wagonRect)
	rgba, ok := screen.(*image.RGBA)
	if !ok {
		return
	}
	wagonMatch := *matching.MatchTemplate(rgba, &templates.WagonTemplate, false)

	for i, wagon := range wagonMatch {
		fmt.Printf("Wagon %d: %s\n", i, wagon.Key)
	}
}

func captureScreen(r rect) image.Image {
	bit := robotgo.CaptureScreen(r.X, r.Y, r.W, r.H)
	defer robotgo.FreeBitmap(bit)
	img := robotgo.ToImage(bit)
	return img
}
