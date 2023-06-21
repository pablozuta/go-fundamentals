package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// open a file for writing
	file, err := os.OpenFile("logFile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error al abrir archivo", err)
	}
	defer file.Close()
	// create a logger that writes to both standard error and the log file
	logger := log.New(io.MultiWriter(os.Stderr, file), "", log.Ldate|log.Ltime)

	// print a message withe the custom logger
	logger.Println("Hola desde custom logger multiple output")
}
