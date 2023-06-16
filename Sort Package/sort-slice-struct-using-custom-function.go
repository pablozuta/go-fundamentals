package main

import (
	"fmt"
	"sort"
)

type Person4 struct {
	Name string
	Age  int
}

func main() {
	people := []Person4{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
		{"David", 30},
		{"Eve", 25},
	}

	// se hace un sort con una funcion que ordena los fields primero por edad (ascendente)
	// y luego por nombre (descendente)
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name > people[j].Name
		} else {
			return people[i].Age < people[j].Age
		}
	})
	// se muestran los resultados
	for _, item := range people {
		fmt.Println(item)
	}
}
