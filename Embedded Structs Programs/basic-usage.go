package main

import "fmt"

type Person struct {
	Name string	
	Age int
}
type Employee struct {
	Person // embedded struct
	Salary float64
}


func main() {
	e := Employee {
		Person: Person{Name: "John Smith", Age: 25},
		Salary: 4500.5,
	}

	// mostrar datos
	fmt.Println(e.Person.Name)
	fmt.Println(e.Salary)
	
}