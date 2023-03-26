package main

import (
	"fmt"
	"strings"
)

func main() {
	// largo de un string
	str := "Die Pyramide"
	fmt.Println(len(str))

	// indexing
	fmt.Println(str[1]) // mostrara el codigo ASCII de la letra

	// rune indexing (mostrara la letra)
	fmt.Println(string(str[0]))
	fmt.Println(string(str[2]))
	fmt.Println(string(str[4]))

	// concatenacion
	str1 := "Gefangen bei "
	str2 := "den Pharaonen"
	fmt.Println(str1 + str2)

	// ver si contiene una palabra
	fmt.Println(strings.Contains(str , "Haus")) // false
	fmt.Println(strings.Contains(str , "Die")) // true

	
}