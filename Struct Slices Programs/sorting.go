package main

import (
	"fmt"
	"sort"
)

// creamos un struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// inicializamos un slice de structs Person
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Kirk", Age: 62},
		{Name: "Bob", Age: 21},
		{Name: "Esther", Age: 45},
		{Name: "Karla", Age: 29},
	}

	// mostramos es slice como fue creado
	fmt.Println("Unsorted", people)

	// usamos el metodo sort.Slice que toma dos argumentos , un slice y una funcion
	// la funcion usa dos indices i y j que cambian de valor automaticamente para hacer la comparacion
	sort.Slice(people, func(i, j int) bool {
		// mostramos los indices i y j solo como aclaracion del algoritmo sort.Slice()
		fmt.Println("Index i", i)
		fmt.Println("Index j", j)
		// ordenamos de menor a mayor 
		return people[i].Name < people[j].Name
	})

	fmt.Println("Sorted by Name", people)
	
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by Age", people)
}
