// This program defines two struct types: Character and Story. The Story struct has a field called Characters which is a slice of Character instances. This is an example of embedding a slice into a struct.
package main

import "fmt"

type character struct {
	name       string
	occupation string
}

type story struct {
	title      string
	characters []character
}

func main() {
	// create a new story instance
	burningChrome := story{
		title: "Burning Chrome",
		characters: []character{
			character{
				name:       "Automatic Jack",
				occupation: "Hacker",
			},
			{
				name:       "Bobby Quine",
				occupation: "Hacker",
			},
			{
				name:       "Rikki",
				occupation: "Stripper",
			},
		},
	}
	// print out the story's information
	fmt.Printf("Title: %s\n", burningChrome.title)
	fmt.Println("Characters:")
	for _, character := range burningChrome.characters {
		fmt.Printf("Name: %s\n", character.name)
		fmt.Printf("Occupation: %s\n", character.occupation)
		fmt.Println()
	}

}
