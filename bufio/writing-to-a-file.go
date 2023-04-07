package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	archivo, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Println("no se pudo crear archivo")
		return
	}

	defer archivo.Close()

	writer := bufio.NewWriter(archivo)
	fmt.Fprintln(writer, "esto se escribira usando bufio")
	fmt.Println("se creo el archivo exitosamente")
	writer.Flush()
}