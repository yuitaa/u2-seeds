package capture

import (
	"fmt"
	"image"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/yuitaa/u2-seeds/internal/matching"
	"github.com/yuitaa/u2-seeds/internal/templates"
)

type Information struct {
	UpperTask      string
	UpperExtension string
	LowerTask      string
	LowerExtension string
	Wagons         []string
}

func (i *Information) String() string {
	return fmt.Sprintf(
		"Task: %s, Extension: %s\nTask: %s, Extension: %s\nWagons: %s",
		i.UpperTask, i.UpperExtension,
		i.LowerTask, i.LowerExtension,
		strings.Join(i.Wagons, ", "),
	)
}

func GetInformation() (*Information, bool) {
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

	var wagons []string
	screen := captureScreen(wagonRect)
	if rgba, ok := screen.(*image.RGBA); ok {
		wagonMatch := *matching.MatchTemplate(rgba, &templates.WagonTemplate, false)
		for _, wagon := range wagonMatch {
			wagons = append(wagons, wagon.Key)
		}
	}

	info := &Information{
		UpperTask:      upperTask,
		UpperExtension: upperExtension,
		LowerTask:      lowerTask,
		LowerExtension: lowerExtension,
		Wagons:         wagons,
	}

	ok := upperTask != "" && upperExtension != "" && lowerTask != "" && lowerExtension != "" && len(wagons) > 0

	return info, ok
}

func captureScreen(r rect) image.Image {
	bit := robotgo.CaptureScreen(r.X, r.Y, r.W, r.H)
	defer robotgo.FreeBitmap(bit)
	img := robotgo.ToImage(bit)
	return img
}
