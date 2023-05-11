package main

import "fmt"

type Person struct {
	Name string
	Age int
}
type Employee struct {
	Person // embed the Person struct
	Job string
	Salary float64
}
type Manager struct {
	Employee // embed the Employee struct
	TeamSize int
}

func main() {
	// create an employee
	e := Employee {
		Person: Person{Name: "Joanne Archs", Age: 23},
		Job: "Software Engineer",
		Salary: 75000.0,
	}

	// create a manager
	m := Manager {
		Employee: Employee{
			Person: Person{Name: "Russ Cox", Age: 34},
			Job:"Lead Architect",
			Salary: 104000.0,
		},
		TeamSize: 10,
	}

	// print info
	fmt.Println(e.Age)
	fmt.Println(e.Job)
	fmt.Println(e.Name)
	fmt.Println(e.Person)
	fmt.Println(e.Salary)

	fmt.Println(m.Age)
	fmt.Println(m.Job)
	fmt.Println(m.Employee)
	fmt.Println(m.Name)
	fmt.Println(m.TeamSize)
	fmt.Println(m.Salary)
	
}