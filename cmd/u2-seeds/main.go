package main

import (
	"github.com/yuitaa/u2-seeds/internal/capture"
	"github.com/yuitaa/u2-seeds/internal/move"
	"github.com/yuitaa/u2-seeds/internal/seed"
)

func main() {
	seedGenerator := seed.NewSeedGenerator()
	move.PressAltTab()

	for i := range 10 {
		nothing(i)
		seed := seedGenerator.Generate()
		move.RestartMove(seed)
		println("Seed: " + seed)
		capture.GetInformation()
		println()
	}
}

func nothing(a int) {}
