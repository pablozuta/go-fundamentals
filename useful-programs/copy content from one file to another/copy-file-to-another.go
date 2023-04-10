package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// definimos los archivos source y destination
	archivoSource := "source.txt"
	archivoDestination := "destination.txt"

	// abrimos el archivo source
	source, err := os.Open(archivoSource)
	if err != nil {
		log.Fatal(err)
	}
	
	defer source.Close()
	
	// creamos el archivo de destino
	destination, err := os.Create(archivoDestination)
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()
	
	// copiar el contenido de un archivo al otro
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	}

	// mensaje despues de realizar la operacion
	fmt.Printf("contenido archivo %s copiado a %s", archivoSource, archivoDestination)

}