package capture

type rect struct {
	X, Y, W, H int
}

var (
	upperTaskRect = rect{X: 0, Y: 174, W: 1300, H: 32}
	lowerTaskRect = rect{X: 0, Y: 211, W: 1300, H: 32}
	wagonRect     = rect{X: 850, Y: 177, W: 1070, H: 32}
)
