package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


func main() {
	p1 := Person{"John", 30}
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data in Json format")	
	fmt.Printf("%s\n", bytes)
	
	var p2 Person
	err = json.Unmarshal(bytes, &p2)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Data in string format")	
	fmt.Printf("%+v\n", p2)
}