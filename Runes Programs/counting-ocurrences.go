package main

import (
	"fmt"
)

func countOccurrences(str string, target rune) int {
	count := 0
	for _, item := range str {
		if item == target {
			count++
		}
	}
	return count
}

func main() {
	str := "Hola Golang"
	target := 'o'
	count := countOccurrences(str, target)
	fmt.Println("Numero de ocurrencias del target: ", count)
}
