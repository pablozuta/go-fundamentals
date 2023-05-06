package main

import (
	"fmt"
	"strings"
)

// creamos un structs para guardar data de los personajes
type character struct {
	name      string
	bookAnime string
	gender    string
	age       int
}

func main() {
	// creamos un array de structs con datos de personajes
	characters := []character{
		{name: "Case", bookAnime: "Neuromancer", gender: "Male", age: 24},
		{name: "Molly Millions", bookAnime: "Neuromancer", gender: "Female", age: 28},
		{name: "Hiro Protagonist", bookAnime: "Snow Crash", gender: "Male", age: 33},
		{name: "Raven", bookAnime: "Snow Crash", gender: "Male", age: 27},
		{name: "Takeshi Kovacs", bookAnime: "Altered Carbon", gender: "Male", age: 42},
		{name: "Quellcrist Falconer", bookAnime: "Altered Carbon", gender: "Female", age: 300},
		{name: "Motoko Kusanagi", bookAnime: "Ghost in the Shell", gender: "Female", age: 28},
		{name: "Batou", bookAnime: "Ghost in the Shell", gender: "Male", age: 38},
	}

	// preguntamos al usuario por un termino de busqueda que puede ser nombre o titulo de anime/libro
	var searchTerm string
	fmt.Print("Enter a character name or book/anime title: ")
	fmt.Scan(&searchTerm)

	// search for characters for name or book/anime
	results := make([]character, 0)
	for _, c := range characters {
		if strings.Contains(strings.ToLower(c.name),
			strings.ToLower(searchTerm)) ||
			strings.Contains(strings.ToLower(c.bookAnime),
				strings.ToLower(searchTerm)) {
			results = append(results, c)
		}
	}

	// Print the search results
	if len(results) == 0 {
		fmt.Println("No results found.")
	} else {
		fmt.Printf("%d results found:\n", len(results))
		for _, c := range results {
			fmt.Printf("- %s (%s, %s, %d years old)\n", c.name, c.bookAnime, c.gender, c.age)
		}
	}

}
