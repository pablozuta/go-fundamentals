package main

import (
	"log"
	"os"
)

func main() {
	// se abre un archivo para escribir
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error al abrir archivo", err)
	}
	defer file.Close()
	// se usa el package log para setear la salida usando el archivo en vez de standard output
	log.SetOutput(file)
	// se crea un mensaje que se grabara en el archivo
	log.Println("Tercer Mensaje de prueba")
}

// This program logs a message to a file instead of standard output. It uses the os.OpenFile function to create or open a file named logfile.txt for writing. The os.O_CREATE flag tells OpenFile to create the file if it doesn't exist, os.O_APPEND tells it to append to the file instead of overwriting it, and os.O_WRONLY tells it to open the file in write-only mode. If OpenFile encounters an error, the program calls log.Fatal() to log an error message and exit.
// After the file is opened, the program sets the log package's output to the file using log.SetOutput(). Any subsequent log messages will be written to the file instead of standard error.
