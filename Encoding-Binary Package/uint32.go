package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, 48241983)

	// decoding back
	decodedBuffer := binary.BigEndian.Uint32(buffer)
    fmt.Println(decodedBuffer)

	fmt.Printf("Encoded Byte Slice: %v", buffer)

}