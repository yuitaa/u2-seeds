package move

import (
	"time"

	"github.com/go-vgo/robotgo"
)

const pauseDuration = 160

func tap(key string, millisecond int) {
	robotgo.KeyTap(key)
	time.Sleep(time.Duration(millisecond) * time.Millisecond)
}
