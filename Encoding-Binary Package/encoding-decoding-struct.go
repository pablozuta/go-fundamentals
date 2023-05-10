package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// encoding struct into byte slice
	buffer := new(bytes.Buffer)
	order := binary.LittleEndian
	err := binary.Write(buffer, order, &MyStruct{42, true})
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// decoding byte slice
	var decoded MyStruct
	err = binary.Read(buffer, order, &decoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print
	fmt.Printf("Encoded byte slice %v", buffer.Bytes())
	fmt.Printf("Decoded struct %+v", decoded) // verbose output of struct
	
}

type MyStruct struct {
	Value uint32
	Flag bool
}