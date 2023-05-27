package main

import (
	"fmt"
	"strings"
)

type Text struct {
	Content string
}

func (t Text) Translate() string {
	translations := map[string]string{
		"Man":       "Mortal",
		"woman":     "maiden",
		"sun":       "burning orb",
		"moon":      "eldritch orb",
		"the night": "the void",
		"the day":   "the waking world",
		"the earth": "the plane of mortal existence",
		"ocean":     "the abyss",
		"death":     "the final slumber",
		"horror":    "unspeakable terror",
	}
	translated := t.Content
	for original, replacement := range translations {
		translated = strings.ReplaceAll(translated, original, replacement)
	}
	return translated
}

func main() {
	paragraph := Text{"Man is but a speck of dust in the grand scheme of things. the sun sets and the moon rises, and the night consumes the day. the earth is but a tiny island in a vast ocean of darkness and horror ."}
	fmt.Println(paragraph)
	fmt.Println()
	fmt.Println(paragraph.Translate())
}
