package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string) int {
	words := strings.Fields(sentence)
	return len(words)
}

func main() {
	sentence := "The quick brown fox jumps over the lazy dog."
	numberOfWords := countWords(sentence)
	fmt.Println("Numero de Palabras en  la frase:", numberOfWords)

}
