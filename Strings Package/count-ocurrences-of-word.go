package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "In a world where consciousness can be transferred to different bodies, Takeshi Kovacs, a former soldier, is hired to solve the murder of a wealthy man." 

	word := "solve"
	wordCount := wordOcurrences(text, word)
	fmt.Printf("The word '%s' occurs %d times in the text", word, wordCount)
}

func wordOcurrences(s, word string)int  {
	// split the text into words
	words := strings.Fields(s)

	// count the ocurrences of the word
	wordCount := 0
	for _, w := range words {
		if strings.EqualFold(w, word) {
			wordCount ++
		}
	}
	return wordCount
}