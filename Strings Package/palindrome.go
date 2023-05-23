package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "level"

	// reverse the word
	reversedWord := reverseString(word)

	// check if the word is palindrome
	isPalindrome := strings.ToLower(word) == strings.ToLower(reversedWord)

	if isPalindrome {
		fmt.Printf("La Palabra '%s' es Palindrome\n", word)
		} else {
		fmt.Println("La Palabra NO es Palindrome")

	}
}

func reverseString(s string)string  {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1;i < j;i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}