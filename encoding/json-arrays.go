package main

import (
	"encoding/json"
	"fmt"
)

type Replicant struct {
	Name         string `json:"name"`
	Model        string `json:"model"`
	InceptDate   string `json:"incept_date"`
	Manufacturer string `json:"manufacturer"`
}

func main() {
	// create an array of replicant objects
	replicants := []Replicant {
		{Name: "Roy Batty", Model: "Nexus-6", InceptDate: "8 Jan 2016", Manufacturer: "Tyrell Corporation"},
		{Name: "Pris Stratton", Model: "Nexus-6", InceptDate: "14 Feb 2016", Manufacturer: "Tyrell Corporation"},
		{Name: "Zhora Salome", Model: "Nexus-6", InceptDate: "12 Jun 2016", Manufacturer: "Tyrell Corporation"},
	}

	// encode the array to JSON
	jsonBytes, err := json.Marshal(replicants)
	if err != nil {
		panic(err)
	}

	// print the JSON string
	fmt.Println(string(jsonBytes))

	// decode the JSON object to an array of Replicant objects
	var replicants2 []Replicant
	err = json.Unmarshal(jsonBytes, &replicants2)
	if err != nil {
		panic(err)
	}

	// print the array of Replicant objects
	fmt.Printf("%+v\n", replicants2)

	// for range to display Replicants
	for _, replicant := range replicants2 {
		fmt.Println("Name:",replicant.Name)
		fmt.Println("Model:",replicant.Model)
		fmt.Println("Manufacturer:",replicant.Manufacturer)
		fmt.Println("--------------------------")
		
	}

}
