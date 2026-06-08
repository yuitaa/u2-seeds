package move

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/yuitaa/u2-seeds/internal/seed"
)

func RestartMove(seed seed.SeedString) {
	// ゲーム退出
	tap(robotgo.Escape, pauseDuration)
	tap(robotgo.KeyS, pauseDuration)
	tap(robotgo.Space, pauseDuration)
	tap(robotgo.KeyA, pauseDuration)
	tap(robotgo.Space, pauseDuration)
	time.Sleep(900 * time.Millisecond)

	// 構成変更
	robotgo.KeyDown(robotgo.KeyW)
	robotgo.KeyDown(robotgo.KeyD)
	time.Sleep(900 * time.Millisecond)
	robotgo.KeyUp(robotgo.KeyW)
	time.Sleep(400 * time.Millisecond)
	robotgo.KeyUp(robotgo.KeyD)
	time.Sleep(800 * time.Millisecond)
	robotgo.KeyDown(robotgo.Space)
	time.Sleep(400 * time.Millisecond)
	robotgo.KeyUp(robotgo.Space)
	time.Sleep(400 * time.Millisecond)

	// シードの入力欄に移動
	tap(robotgo.KeyS, pauseDuration)
	tap(robotgo.KeyS, pauseDuration)
	tap(robotgo.KeyS, pauseDuration)
	tap(robotgo.Space, pauseDuration)

	// シード入力
	for _, char := range string(seed) {
		tap(string(char), pauseDuration)
	}
	tap(robotgo.Space, pauseDuration)
	tap(robotgo.Escape, pauseDuration)

	// ゲーム開始
	robotgo.KeyDown(robotgo.KeyS)
	robotgo.KeyDown(robotgo.KeyA)
	time.Sleep(1200 * time.Millisecond)
	robotgo.KeyUp(robotgo.KeyS)
	time.Sleep(1900 * time.Millisecond)
	robotgo.KeyUp(robotgo.KeyA)

	time.Sleep(8000 * time.Millisecond)
	tap(robotgo.KeyT, pauseDuration)
	time.Sleep(18000 * time.Millisecond)

}
