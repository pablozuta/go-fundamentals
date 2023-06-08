package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string {"apple", "banana", "cherry", "date", "elderberry"}

	// sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	
	fmt.Println("Ordered words by lenght:", words)
}