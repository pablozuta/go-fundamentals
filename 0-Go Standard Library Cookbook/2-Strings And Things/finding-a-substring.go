package main

import (
	"fmt"
	"strings"
)

const refString = "Mery had a little lamb"

func main() {
	lookFor := "lamb"
	// el methodo Contains toma dos argumentos 1-string donde se buscara substring 2
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	startsWith := "Mery"
	// el metodo HasPrefix toma dos argumentos 1-string en la que buscar(primera palabra)
	// 2- palabra que se buscara
	starts := strings.HasPrefix(refString, startsWith)
	fmt.Printf("The \"%s\" starts with \"%s\": %t \n", refString, startsWith, starts)

	endsWith := "lamb"
	// el metodo HasSuffix toma dos argumentos 1-string en la que buscar(ultima palabra)
	// 2- palabra que se buscara
	ends := strings.HasSuffix(refString, endsWith)
	fmt.Printf("The \"%s\" ends with \"%s\": %t \n", refString, endsWith, ends)

}
