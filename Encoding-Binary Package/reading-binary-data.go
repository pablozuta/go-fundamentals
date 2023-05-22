package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	// abrimos el archivo para leer
	file, err := os.Open("data.bin")
	if err != nil {
		fmt.Println(err)
		return
	} 
	defer file.Close()

	var value uint32
	err = binary.Read(file, binary.LittleEndian, &value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Value read from file: %d", value)

	
}