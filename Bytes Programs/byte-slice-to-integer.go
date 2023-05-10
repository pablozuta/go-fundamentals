package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	bytes := []byte{0x01, 0x02, 0x03, 0x03}
	num := binary.LittleEndian.Uint32(bytes)
	fmt.Println(num)
}