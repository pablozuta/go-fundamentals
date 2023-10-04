// In this example, we have a character struct and an organization struct that both embed a city struct. The city struct contains information about a city in the anime, while the character struct contains information about a character and the organization struct contains information about an organization in the anime.

package main

import "fmt"

type city struct {
	name    string
	country string
	size    string
}

type character struct {
	city    // embed the city struct
	name    string
	age     int
	company string
}

type organization struct {
	city    // embed the city struct
	name    string
	members []string
}

func main() {
	// create a city
	c := city{
		name:    "New Port City",
		country: "Japan",
		size:    "Large",
	}

	// create a character
	ch := character{
		city:    c, // embed the city struct
		name:    "Motoko Kusagani",
		age:     37,
		company: "Public Security Section 9",
	}

	// create an organization
	o := organization{
		city:    c, // embed the city struct
		name:    "Public Security Section 9",
		members: []string{"Batou", "Togusa", "Ishikawa", "Saito", "Pazu", "Borma"},
	}

	// print some information about the character
	fmt.Println("Name:", ch.name)
	fmt.Println("Age:", ch.age)
	fmt.Println("Company:", ch.company)
	fmt.Println("City:", ch.city.name)
	fmt.Println("Country:", ch.city.country)
	fmt.Println("Size:", ch.city.size)

	// print some information about the organization
	fmt.Println("Name:", o.name)
	fmt.Println("Members:", o.members)
	fmt.Println("City:", o.city.name)
	fmt.Println("Country:", o.city.country)
	fmt.Println("Size:", o.city.size)
}
