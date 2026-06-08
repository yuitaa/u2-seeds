package templates

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

var (
	BiomeTemplate, _     = loadTemplates("biomes")
	TaskTemplate, _      = loadTemplates("tasks")
	ExtensionTemplate, _ = loadTemplates("extensions")
	WagonTemplate, _     = loadTemplates("wagons")
)

type TemplateMap = map[string]*image.RGBA

func loadTemplates(dir string) (TemplateMap, error) {
	imageMap := make(TemplateMap)
	templateDir := filepath.Join("internal/templates", dir)

	files, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if ext := filepath.Ext(file.Name()); strings.ToLower(ext) != ".png" {
			continue
		}

		path := filepath.Join(templateDir, file.Name())
		f, err := os.Open(path)
		if err != nil {
			continue
		}
		defer f.Close()

		template, _, err := image.Decode(f)
		if err != nil {
			continue
		}

		bounds := template.Bounds()
		rgbaImg := image.NewRGBA(bounds)

		draw.Draw(rgbaImg, bounds, template, bounds.Min, draw.Src)

		imageMap[getFileName(file.Name())] = rgbaImg
	}
	return imageMap, nil
}

func getFileName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}
