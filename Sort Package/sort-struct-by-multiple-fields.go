package main

import (
	"fmt"
	"sort"
)

type Person2 struct {
	Name string
	Age  int
}

func main() {
	people := []Person2{
		{"Yann", 57},
		{"John", 37},
		{"Charlie", 30},
		{"Fabby", 87},
		{"Elisa", 30},
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name > people[j].Name
		} else {
			return people[i].Age < people[j].Age
		}
	})

	fmt.Println(people)
}

// In this example, we're sorting a slice of Person structs by multiple fields using the sort.Slice() function. We pass a custom comparison function to sort.Slice() that first compares the ages of the people, and then the names if the ages are the same. If the ages are the same, we sort by name in descending order by using the greater than (>) operator instead of the less than (<) operator.
