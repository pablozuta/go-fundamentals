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

	// contar la repeticion de cierta letra
	fmt.Println(strings.Count(str, "n"))

	// split at character
	var frase = "Una.Frase.Unida.Por.Puntos"
	fmt.Println(strings.Split(frase, "."))

	// to upper case
	fmt.Println(strings.ToUpper(str1))

	// to lower
	fmt.Println(strings.ToLower(str1))

	// replace
	var saludo = "Hello Golang"
	fmt.Println(saludo)
	fmt.Println(strings.Replace(saludo, "Hello", "Hi", 1))


	
}