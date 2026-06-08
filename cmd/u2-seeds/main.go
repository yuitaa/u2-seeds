package main

import (
	"fmt"

	"github.com/yuitaa/u2-seeds/internal/templates"
)

func main() {
	templ := templates.WagonTemplate
	for key, value := range templ {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
