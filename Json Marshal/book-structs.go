package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func main() {
	book := Book{
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
		Price:  12.99,
	}
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(jsonData))
}
