package main

import (
	"encoding/json"
	"fmt"
)

type Cyberpunk struct {
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Location   string `json:"location"`
	Cyberware  struct {
		AugmentedEyes bool   `json:"augmented_eyes"`
		NeuralImplant string `json:"neural_implant"`
	} `json:"cyberware"`
}


func main() {
	// create a cyberpunk object
	var cp Cyberpunk
	cp.Name = "Johnny Silverhand"
	cp.Profession = "Mercenary"
	cp.Location = "Night City"
	cp.Cyberware.AugmentedEyes = true
	cp.Cyberware.NeuralImplant = "Chippin In"

	// encode the object to JSON
	jsonBytes, err := json.Marshal(cp)
	if err != nil {
		panic(err)
	}

	// print the JSON string
	fmt.Println("JSON string")
	fmt.Println(string(jsonBytes))
	fmt.Println()

	// decode the JSON string to a Cyberpunk object
	var cp2 Cyberpunk
	err = json.Unmarshal(jsonBytes, &cp2)
	if err != nil {
		panic(err)
	}

	// print the Cyberpunk object
	fmt.Println("Cyberpunk Object")
	fmt.Printf("%+v", cp2)
}
