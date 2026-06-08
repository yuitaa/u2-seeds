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
		var bestMatch *ScoreEntry

		for dx := range imgW - templateW {
			score, _ := calculateScore(img, template, dx)

			if score > threshold {
				if bestMatch == nil || dx > bestMatch.Dx+10 {
					if bestMatch != nil {
						results = append(results, *bestMatch)
					}
					bestMatch = &ScoreEntry{Key: key, Score: score, Dx: dx}
				} else if score > bestMatch.Score {
					bestMatch.Score = score
					bestMatch.Dx = dx
				}

				if matchOnce {
					results = append(results, *bestMatch)
					break TemplateLoop
				}
			}
		}
		if bestMatch != nil && !matchOnce {
			results = append(results, *bestMatch)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Dx < results[j].Dx
	})

	return &results
}
