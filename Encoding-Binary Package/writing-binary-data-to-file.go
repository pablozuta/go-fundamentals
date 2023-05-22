package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	// creamos archivo
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var value uint16 = 12345
	err = binary.Write(file, binary.LittleEndian, &value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Valor escrito a archivo exitosamente.")
}