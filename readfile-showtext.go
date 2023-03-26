package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// leer el contenido del archivo y guardarlo en una variable de tipo slice
	fileContent, err := ioutil.ReadFile("output.txt")
	if err != nil {
		panic(err)
	}

	// convertir variable slice a string
	fileContentString := string(fileContent)

	// mostrar el contenido por consola
	fmt.Println(fileContentString)
}