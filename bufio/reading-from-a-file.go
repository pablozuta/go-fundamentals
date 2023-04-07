package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("texto.txt")
	if err != nil {
		fmt.Println("No se pudo leer el archivo", err)
		return

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al scanear archivo", err)


	}
}
