package capture

type Rect struct {
	X, Y, W, H int
}

var (
	UpperTaskRect = Rect{X: 64, Y: 174, W: 1200, H: 32}
	LowerTaskRect = Rect{X: 64, Y: 211, W: 1200, H: 32}
	WagonRect     = Rect{X: 850, Y: 177, W: 1070, H: 32}
)
