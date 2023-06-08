package main

import (
	"fmt"
	"sort"
)

// se crea struct de tipo Person
type Person struct {
	Name string
	Age int
}

func main() {
	// se crea slice de objetos de tipo Person
	people := []Person {
		{"Alice", 23},
		{"John", 33},
		{"Sofia", 27},
	}
	// se muestra el slice 
	fmt.Println(people)

	// se usa funcion sort.Slice usando el field Age 
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	// se muestra el resultado ordenado
	fmt.Println(people)
}