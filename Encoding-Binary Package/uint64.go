package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(buffer, 483728134124)
	fmt.Printf("Encoded Byte Slice %v\n", buffer)

	value := binary.LittleEndian.Uint64(buffer)
	fmt.Println("Decoded Byte Slice" , value)

}