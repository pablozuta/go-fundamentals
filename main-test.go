package main

import "fmt"

func main() {
	bytes := []byte{'h','o', 'l', 'a'}
	fmt.Println(len(bytes))
	fmt.Println(bytes)

	saludo := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	saludoString := string(saludo)
	fmt.Println(saludoString)
}