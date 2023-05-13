package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	// if the list is empty return an empty string
	if len(strs) == 0 {
		return ""
	}
	// define a 2D array to store the characters of each string
	chars := make([][]byte, len(strs))
	for i := range chars {
		chars[i] = []byte(strs[i])
	}

	// find the length of the shortest string
	shortestLen := len(chars[0])
	for i := 1; i < len(chars); i++ {
		if len(chars[i]) < shortestLen {
			shortestLen = len(chars[i])
		}
	}

	// compare each character in the strings
	for i := 0; i < shortestLen; i++ {
		for j := 1; j < len(chars); j++ {
			if chars[j][i] != chars[0][i] {
				return string(chars[0][:i])
			}
		}
	}

	// if all characters match return the first string
	return string(chars[0][:shortestLen])
}

func main() {
	strs := []string{"neuromancer", "neuroskull", "neurofunk"}
	fmt.Println(longestCommonPrefix(strs))
}
