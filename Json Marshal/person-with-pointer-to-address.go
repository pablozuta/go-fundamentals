package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
}
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

func main() {
	person := Persona{
		Name: "John Smith",
		Age:  30,
		Address: &Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			ZipCode: "12345",
		},
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))

}
