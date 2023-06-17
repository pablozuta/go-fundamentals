package main

import (
	"encoding/json"
	"fmt"
)

type User2 struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {
	people := []User2{
		{"Alice", 32, "alice@gmail.com", "52 london street"},
		{"Paul", 34, "paul32@gmail.com", "Grand Avenue 3423"},
	}
	// store the data as json
	jsonData, _ := json.Marshal(people)

	// retrieve the data
	var retrievedUsers []User2
	err := json.Unmarshal(jsonData, &retrievedUsers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print the retrieved data
	fmt.Println(retrievedUsers)
}
