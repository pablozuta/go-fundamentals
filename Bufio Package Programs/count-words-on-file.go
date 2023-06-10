package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error al abrir archivo", err)
		return
	}
	defer file.Close()

	// create a map to store word frequencies
	wordCount := make(map[string]int)

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// set the scanner to split on words
	scanner.Split(bufio.ScanWords)

	// iterate over each word on the file
	for scanner.Scan() {
		// pasamos las palabra a lowercase
		word := strings.ToLower(scanner.Text())
		// agregamos la palabra al contador(mapa) como key y la cuenta como valor integer
		wordCount[word]++
	}
	// check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scaneando archivo", err)
		return
	}

	// mostramos la palabra y la frecuencia iterando en los elementos del mapa
	for word, count := range wordCount {
		fmt.Printf("palabra :%s   numero de ocurrencias:%d\n", word, count)
	}
}