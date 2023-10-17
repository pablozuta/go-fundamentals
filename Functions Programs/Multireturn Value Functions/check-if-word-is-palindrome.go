package main

import "fmt"

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := "level"
	result := isPalindrome(input)

	if result {
		fmt.Printf("Yes, word '%s' is Palindrome\n", input)
	} else {
		fmt.Println("Word is NOT Palindrome")
	}
}
