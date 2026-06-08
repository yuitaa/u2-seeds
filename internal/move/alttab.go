package move

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func PressAltTab() {
	robotgo.KeyDown(robotgo.Alt)
	robotgo.KeyTap(robotgo.Tab)
	robotgo.KeyUp(robotgo.Alt)
	time.Sleep(pauseDuration * time.Millisecond)
}
