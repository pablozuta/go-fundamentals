package main

import "fmt"

func main() {
	// initialize a map
	myMap := make(map[string]int)
		myMap["uno"] = 1
		myMap["dos"] = 2
		myMap["tres"] = 3
	// print the map
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	// otra manera de inicializar un mapa
	coltrane := map[int]string {
			1957: "Blue Train",
			1959: "Giant Steps",
			1960: "My Favorite Things",
			1961: "Ole Coltrane",
			1963: "Ballads",
		    1964: "A Love Supreme",
	}
	// print the map
	for year, title := range coltrane {
		fmt.Printf("%d: %s\n", year, title)
	}

	// imprimiendo valores usando keys
	fmt.Println(coltrane[1960])

}