package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "the quick brown fox"

	// capitalize first letter of each word
	capitalizedSentence := capitalizeWords(sentence) 

	fmt.Println("Original Sentence:", sentence)
	fmt.Println("Capitalized Sentence:", capitalizedSentence)
}

func capitalizeWords(s string)string  {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}