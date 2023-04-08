package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	jsonString := `{"Name": "Bob", "Age": 30}`
	var persona Person
	err := json.Unmarshal([]byte(jsonString), &persona)

	if err != nil {
		fmt.Println("no se pudo realizar decoding")
	}
	fmt.Println("Decoding da como resultado el siguiente struct:")
	fmt.Printf("%+v", persona)


}