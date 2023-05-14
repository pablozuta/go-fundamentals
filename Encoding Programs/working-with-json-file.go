package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type Drink struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

type Menu struct {
	Drinks []Drink `json:"drinks"`
}


func main() {
	// read the json file
	fileBytes, err := os.ReadFile("menu.json")
	if err != nil {
		fmt.Println("No se pudo encontrar el archivo especificado")
		panic(err)
	}

	// decode the JSON file to a Menu object
	var menu Menu
	err = json.Unmarshal(fileBytes, &menu)
	if err != nil {
		panic(err)
	}

	// mostrar el nombre y costo del primer objeto del array menu
	fmt.Printf("%s - Price $ %d", menu.Drinks[0].Name, menu.Drinks[0].Price)
	
}