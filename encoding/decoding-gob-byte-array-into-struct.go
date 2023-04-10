package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	p.Name = "Alice"
	p.Age = 30

	// Encode the person struct
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	// Decode the encoded byte slice
	var decodedPerson Person
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&decodedPerson)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Printf("%+v\n", decodedPerson)
}