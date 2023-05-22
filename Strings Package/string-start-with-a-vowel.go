package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Neuromancer"

	vowels := []string{"a", "e", "i", "o", "u"}

	// comprobar si la palabra empieza con vocal
	startsWithVowel := false
	for _, vowel := range vowels {
		if strings.HasPrefix(strings.ToLower(word), vowel) {
			startsWithVowel = true
			break
		}
	}
	if startsWithVowel {
		fmt.Printf("The word %s start with a vowel.\n", word)
	} else {
		fmt.Printf("The word %s DO NOT start with a vowel.\n", word)

	}
}
