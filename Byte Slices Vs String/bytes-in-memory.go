package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "Hello, 世界!" // 14 bytes
	str2 := "Hello, World!" // 13 bytes

	length1 := len(str1)
	length2 := len(str2)

	runeCount1 := utf8.RuneCountInString(str1)
	runeCount2 := utf8.RuneCountInString(str2)

	fmt.Println("String 1 length (bytes):",length1)
	fmt.Println("String 2 length (bytes):",length2)
	fmt.Println("String 1 rune count:",runeCount1)
	fmt.Println("String 2 rune count:",runeCount2)
}