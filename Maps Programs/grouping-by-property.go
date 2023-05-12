package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	State string
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 28, State: "California"},
		{Name: "Miles", Age: 38, State: "South Carolina"},
		{Name: "John", Age: 37, State: "South Carolina"},
		{Name: "Pharoah", Age: 67, State: "Lousiana"},
	}

	peopleByState := make(map[string][]Person)

	for _, person := range people {
		state := person.State
		peopleByState[state] = append(peopleByState[state], person)
	}
	fmt.Println(peopleByState)
}
