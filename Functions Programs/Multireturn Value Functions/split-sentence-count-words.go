package main

import (
	"fmt"
	"strings"
)

// la funcion devolvera las palabras por separado y el numero de palabras
func splitAndCount(frase string) ([]string, int) {
	palabras := strings.Fields(frase)
	return palabras, len(palabras)
}

func main() {
	frase := "The quick brown fox jumps over the lazy black cat"
	palabras, numeroPalabras := splitAndCount(frase)

	fmt.Printf("Sentence:%s \nNumber of Words: %d\n", frase, numeroPalabras)
	for _, palabra := range palabras {
		fmt.Println("-", palabra)
	}
}
