package matching

import (
	"image"
	"sort"

	"github.com/yuitaa/u2-seeds/internal/templates"
)

type ScoreEntry struct {
	Key   string
	Score float64
	Dx    int
}

const threshold float64 = 0.96

func MatchTemplate(img *image.RGBA, templates *templates.TemplateMap, matchOnce bool) *[]ScoreEntry {
	var results []ScoreEntry
	imgW := img.Bounds().Dx()

TemplateLoop:
	for key, template := range *templates {
		templateW := template.Bounds().Dx()

		for dx := range imgW - templateW {
			score, _ := calculateScore(img, template, dx)

			if score > threshold {
				results = append(results, ScoreEntry{
					Key:   key,
					Score: score,
					Dx:    dx,
				})
				if matchOnce {
					break TemplateLoop
				}
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Dx < results[j].Dx
	})

	return &results
}
