package main

import (
	"encoding/json"
	"fmt"
)

type Hacker struct {
	Alias      string   `json:"alias"`
	RealName   string   `json:"real_name"`
	Occupation string   `json:"occupation"`
	Skills     []string `json:"skills"`
}

func main() {
	// create an array of hacker objects
	hackers := []Hacker{
		{Alias: "Neo", RealName: "Thomas Anderson", Occupation: "Hacker", Skills: []string{"Programming", "Hacking", "Martial Arts"}},
		{Alias: "Trinity", RealName: "Trinity", Occupation: "Hacker", Skills: []string{"Piloting", "Hacking", "Martial Arts"}},
		{Alias: "Morpheus", RealName: "Morpheus", Occupation: "Resistance Leader", Skills: []string{"Leadership", "Combat", "Philosophy"}},
	}

	// encode the array to JSON
	jsonBytes, err := json.Marshal(hackers)
	if err != nil {
		panic(err)
	}

	// show the JSON string
	fmt.Println("JSON string")
	fmt.Println(string(jsonBytes))
	fmt.Println()

	// decode the JSON string to a Hacker array of objects
	var hackers2 []Hacker
	err = json.Unmarshal(jsonBytes, &hackers2)
	if err != nil {
		panic(err)
	}

	// print the array of Hacker objects
	fmt.Println("Array of Hacker objects")
	fmt.Printf("%+v\n", hackers2)
	fmt.Println()

	// print first array object
	fmt.Printf("%s\n", hackers2[0].Alias)
	fmt.Printf("%s\n", hackers2[0].RealName)
	fmt.Printf("%s\n", hackers2[0].Occupation)
	fmt.Printf("%s\n", hackers2[0].Skills[0])

}
