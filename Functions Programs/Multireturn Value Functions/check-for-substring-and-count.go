package main

import (
	"fmt"
	"strings"
)

// la funcion toma dos strings como argumentos y devuelve dos valores
func checkSubstringCount(str, substr string) (bool, int) {
	// guardamos el numero de veces que la substring se encuentra en la string
	count := strings.Count(str, substr)
	// devolvemos un valor true si el numero count es mayor a 0 y tambien el numero de veces de la substring
	return count > 0, count
}

func main() {
	sentence := "The quick brown fox jumps over the lazy dog. The dog is very lazy."
	substring := "lazy"
	present, count := checkSubstringCount(sentence, substring)
	fmt.Printf("Sentence %s\nSubstring: %s\nPresent = %t\nCount = %d", sentence, substring, present, count)

}
