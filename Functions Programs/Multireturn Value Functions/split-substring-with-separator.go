package main

import (
	"fmt"
	"strings"
)

func splitStringBySeparator(input, separator string) []string {
	return strings.Split(input, separator)
}

func main() {
	sentence := "Yeah bro, RISC arquitecture gonna change everything"
	separator := ","
	substrings := splitStringBySeparator(sentence, separator)

	fmt.Printf("Original string: %s\n", sentence)
	for _, palabra := range substrings {
		fmt.Println("-", palabra)
	}
}
