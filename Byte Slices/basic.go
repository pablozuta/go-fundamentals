package main

import "fmt"

func main() {
	// crear un byte slice
	data := []byte{'H', 'e', 'l', 'l', 'o'}
	// print the slice
	fmt.Println(data)
	// modificar un elemento
	data[0] = 'C'
	fmt.Println(data)
	// append a new element
	data = append(data, '!')
	// convertir a string y mostrar
	fmt.Println(string(data))
}