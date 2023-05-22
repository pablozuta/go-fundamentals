package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "The quick brown fox jumps over the lazy dog."  

	// reemplazar fox with cat
	newText := strings.Replace(text, "fox", "cat", -1)

	fmt.Println("String Original:", text)
	fmt.Println("String con palabra reemplazada:", newText)
}