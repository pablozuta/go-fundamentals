package main

import (
	"bytes"
	"fmt"
)

func main() {
	// crear un buffer y escribir data
	buffer := new(bytes.Buffer)
	buffer.WriteString("Hello")
	buffer.WriteString(" from golang")

	// convertir el buffer a un byte slice y mostrarlo por pantalla
	data := buffer.Bytes()
	fmt.Println(string(data))
}