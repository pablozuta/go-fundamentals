package main

import (
	"log"
	"os"
)

func main() {
	// creamos a new logger con custom prefixes and flags
	logger := log.New(os.Stderr, "myapp: ", log.Ldate|log.Ltime|log.Lshortfile)

	// mostramos un mensaje usando el custom logger
	logger.Println("Hola Logger!!")
}
