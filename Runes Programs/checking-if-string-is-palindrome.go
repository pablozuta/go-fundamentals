package main

import "fmt"

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	str := "level"
	isPal := isPalindrome(str)
	if isPal == true {
		fmt.Println("Frase es palindrome")
	} else {
		fmt.Println("Frase NO es palindrome")

	}

}
