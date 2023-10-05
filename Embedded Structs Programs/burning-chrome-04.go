// This program defines two struct types: place and scene. The scene struct embeds the place struct using the syntax place. This means that the fields of the place struct can be accessed directly from the scene struct.

// In the main function, a new scene instance is created with a description, a place name of "Night City", and a location of "California". The program then prints out the scene's information using fmt.Printf.

package main

import "fmt"

type place struct {
	name     string
	location string
}

type scene struct {
	description string
	place
}

func main() {
	// create a new scene instance
	sceneOne := scene{
		description: "The neon lights reflected off the slick pavement as Automatic Jack and Bobby Quine made their way down the alley.",
		place: place{
			name:     "Night City",
			location: "California",
		},
	}

	// print out the scene information
	fmt.Printf("Description: %s\n", sceneOne.description)
	fmt.Printf("Place Name: %s\n", sceneOne.name)
	fmt.Printf("Location: %s\n", sceneOne.location)
}
