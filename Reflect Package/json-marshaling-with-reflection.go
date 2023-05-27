package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Record struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

func main() {
	// Creating an instance of the Record struct
	p := Record{"A Love Supreme", "John Coltrane", 1964}
	// Getting the reflect.Value of the struct instance
	v := reflect.ValueOf(p)

	// Creating a map to store the JSON data
	jsonData := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		// Getting the reflect.Type of the field
		field := v.Type().Field(i)
		// Getting the value of the JSON tag
		jsonTag := field.Tag.Get("json")
		// Storing the field value in the map using the JSON tag as the key
		jsonData[jsonTag] = v.Field(i).Interface()
	}
	// Marshaling the map into JSON bytes
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error Marshaling JSON", err)
		return
	}
	// Printing the JSON data as a string
	fmt.Println(string(jsonBytes))
}
