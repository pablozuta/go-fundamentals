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
	log.Println("Mensaje de prueba")
}
