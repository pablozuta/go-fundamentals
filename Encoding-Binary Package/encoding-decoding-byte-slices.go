package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// encoding byte slice with little-endian byte order
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, 1234567890)

	// decoding byte slice with big-endian byte order
	value := binary.BigEndian.Uint32(buffer)

	fmt.Printf("Encoded byte slice: %v\n", buffer)
	fmt.Printf("Decoded value: %d\n", value)

}