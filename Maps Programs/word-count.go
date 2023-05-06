package main

import (
	"fmt"
	"strings"
)

func main()  {
	// inicializamos una variable de tipo string
	text := "The cyborg traveled through the wormhole to explore the alien planet The planet was inhabited by a race of telekinetic beings who possessed advanced technology The cyborg was impressed by the beings ability to manipulate matter with their minds He asked them to teach him their ways They agreed"


	// strings.Fields() separa palabras en los espacios, tabs y new lines
	words := strings.Fields(text)

	// creamos un mapa con key string que guardara cada palabra y un valor int para guardar cada ocurrencia de una palabra
	wordCount := make(map[string]int)

	// toma las palabras y las usa como keys en el mapa
	for _, word := range words {
		// si la palabra existe suma 1 al valor
		wordCount[word]++
	}

	// for loop para imprimir las palabras y numero de ocurrencias
	for word, count := range wordCount {
		fmt.Printf("%s %d\n", word, count)
	}
}