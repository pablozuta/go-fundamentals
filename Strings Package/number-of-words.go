package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "The things about blogs is they rarely only talk about go and they often go outdated because good quality post takes lot of time and efforts but here is a list I gathered over the years"

	// separar la frase en los espacios
	words := strings.Split(text, " ")
	// contar el numero de palabras
	wordCount := len(words)

	fmt.Printf("El numero de palabras en la string es %d", wordCount)
}