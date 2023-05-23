package main

import (
	"fmt"
)

func main() {
	text := "jorge luis borges"

	// reverse the string
	reversedText := reverseString(text)
	fmt.Println("Original String:", text)
	fmt.Println("Reversed String:", reversedText)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
