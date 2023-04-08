package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)
type Person struct {
	Name string
	Age int
}

func main() {
	p := Person{"Alice Coltrane", 36}
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)

	if err != nil {
		fmt.Println("no se pudo hacer encoding")
	}
	fmt.Println(buf.Bytes())

}