package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// create a map[string] interface {} to hold data
	data := make(map[string]interface{})
	data["title"] = "Blade Runner"
	data["year"] = 1982
	data["director"] = "Ridley Scott"
	data["actors"] = []string{"Harrison Ford", "Rutger Hauer", "Sean Young"}
	data["characters"] = []map[string]string{
		{"name": "Rick Deckard", "actor": "Harrison Ford"},
		{"name": "Roy Batty", "actor": "Rutger Hauer"},
		{"name": "Pris Stratton", "actor": "Daryl Hannah"},
	}

	// encode the map to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// print the JSON string
	fmt.Println("JSON string -----------------------------------------------------------")
	fmt.Println(string(jsonBytes))

	// decode the JSON string to a map[string] interface{}
	var data2 map[string]interface{}
	err = json.Unmarshal(jsonBytes, &data2)
	if err != nil {
		panic(err)
	}

	// print the map[string] interface{}
	fmt.Println("After Unmarshal --------------------------------------------------------")
	fmt.Printf("%+v", data2)

}
