package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// funcion que recibe un archivo y un channel de tipo int
func wordCount(file *os.File, result chan int)  {
	// se usa un scanner bufio para leer el contenido de file
	scanner := bufio.NewScanner(file)
	// se inicia una variable para guardar el resultado de la cantidad de palabras
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += len(strings.Fields(line))
	}
	result <- count
}

func main() {
	// se usa os para abrir el archivo
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
    
	// se crea un channel que recibira los resultados de count
	result := make(chan int)
	// se inicia una goroutine
	go wordCount(file, result)
	count := <- result
	// se muestra el resultado
	fmt.Printf("Word count = %d", count)
}
