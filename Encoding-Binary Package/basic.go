// This code creates a byte slice b of length 2 and encodes the uint16 value 12345 into it using the little-endian byte order. The fmt.Println statement outputs the byte slice as [57 48], which are the ASCII codes for the characters '9' and '0'. This is because the binary representation of the uint16 value 12345 in little-endian byte order is [57 48].  
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var b []byte = make([]byte, 2)
	binary.LittleEndian.PutUint16(b, 12345)
	fmt.Println(b)

}
