package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// encoding uint16 en slice de bytes de longitud 2
	buffer := make([]byte, 2)
	fmt.Println(buffer) // se inicializa al zero-value
	binary.LittleEndian.PutUint16(buffer, 45678)
	
	// decoding slice de vuelta a uint16
	value := binary.LittleEndian.Uint16(buffer)
	fmt.Println(value) // 45678

	fmt.Printf("Encoded slice %v", buffer)
}