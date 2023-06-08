package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "In a world where consciousness can be transferred to different bodies, Takeshi Kovacs, a former soldier, is hired to solve the murder of a wealthy man."  

	// extract unique words from the text
	uniqueWords := extractUniqueWords(text)

	fmt.Println("Unique words on the text:")
	for _, word := range uniqueWords {
		fmt.Println(word)
	}
}

func extractUniqueWords(s string) []string {
	words := strings.Fields(s)

	// create a map to store unique words
	uniqueWordMap := make(map[string] bool)
	for _, w := range words {
		uniqueWordMap[strings.ToLower(w)] = true
	}

	// extract the unique words from the map
	var uniqueWords []string
	for word := range uniqueWordMap {
		uniqueWords = append(uniqueWords, word)
	}
	return uniqueWords
}