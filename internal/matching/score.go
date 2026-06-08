package matching

import "image"

func calculateScore(img, template *image.RGBA, dx int) (float64, int) {
	imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()
	templateW, templateH := template.Bounds().Dx(), template.Bounds().Dy()

	if imgH != templateH || dx < 0 || dx+templateW > imgW {
		return 0, 0
	}

	maxPossibleDiff := float64(3 * 255 * templateW * templateH)

	imgStride := img.Stride
	templateStride := template.Stride
	imgPix := img.Pix
	templatePix := template.Pix

	rawScore := 0
	for y := range imgH {
		imgRowOff := y*imgStride + dx*4
		templateRowOff := y * templateStride
		for tx := range templateW {
			imgOff := imgRowOff + tx*4
			templateOff := templateRowOff + tx*4
			dr := absInt(int(imgPix[imgOff]) - int(templatePix[templateOff]))
			dg := absInt(int(imgPix[imgOff+1]) - int(templatePix[templateOff+1]))
			db := absInt(int(imgPix[imgOff+2]) - int(templatePix[templateOff+2]))
			rawScore += dr + dg + db
		}
	}
	normalized := 1.0 - (float64(rawScore) / maxPossibleDiff)
	return normalized, rawScore
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
