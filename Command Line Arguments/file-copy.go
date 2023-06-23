package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	// os.Args is a slice of strings that holds all the arguments passed to the program
	// The first argument is the name of the executable itself
	// The second argument is the name of the source file
	// The third argument is the name of the destination file
	args := os.Args
	// abrir el archivo para leer(usando command line arguments)
	source, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer source.Close()
	// crear el archivo destino para escribir
	destination, err := os.Create(args[2])
	if err != nil {
		panic(err)
	}
	defer destination.Close()
	// copiar el contenido de un archivo a otro
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
	fmt.Println("archivo creado y contenido copiado exitosamente")
}
