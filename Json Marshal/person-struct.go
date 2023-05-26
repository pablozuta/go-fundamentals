package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"nombre"`
	Age  int    `json:"edad"`
}

func main() {
	p := Person {"John", 34}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(b))
}
