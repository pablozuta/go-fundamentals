package main

import (
	"bytes"
	"fmt"
)

func main() {
	// hello word en bytes
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}

	// replace hello with go
	newData := bytes.Replace(data, []byte("world"), []byte("golang"), -1)

	fmt.Println(data) // muestra bytes sin formatear (numeros)
	fmt.Println(string(data))
	fmt.Println(string(newData))
}