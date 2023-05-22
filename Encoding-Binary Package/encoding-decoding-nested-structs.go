package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// encode nested struct en un byte slice
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, &OuterStruct{InnerStruct: InnerStruct{42}})
	if err != nil {
		fmt.Println(err)
		return
	}
	// decodificar byte slice de vuelta a nested struct
	var decoded OuterStruct
	err = binary.Read(buffer, binary.LittleEndian, &decoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encoded byte slice: %v\n", buffer.Bytes())
	fmt.Printf("Decoded Nested Struct: %+v\n", decoded)

}

type InnerStruct struct {
	Value uint32
}

type OuterStruct struct {
	InnerStruct
	Flag bool
}
