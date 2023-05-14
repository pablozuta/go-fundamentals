package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func main() {
	libro := Book{"Pedro Paramo", "Juan Rulfo", 1956, "Realismo Magico"}
	l, err := json.Marshal(libro)
	if err != nil {
		fmt.Println("No se pudo hacer encoding")
	}
	fmt.Println("Go struct title: ", libro.Title)
	fmt.Println("Go struct author: ", libro.Author)
	fmt.Println("Go struct year: ", libro.Year)
	fmt.Println("Go struct genre: ", libro.Genre)
	fmt.Println("Json: ", string(l))
}
