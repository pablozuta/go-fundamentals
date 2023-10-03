// In this example, we have an employee struct and a manager struct that both embed a person struct. The person struct contains basic information about the person, while the employee struct contains additional information about their job and salary, and the manager struct contains additional information about their team.
package main

import "fmt"

type person struct {
	name string
	age  int
}
type employee struct {
	person // embed the person struct
	job    string
	salary float64
}
type manager struct {
	employee // embed the employee struct
	teamSize int
}

func main() {
	// create an employee
	e := employee{
		person: person{name: "Joanne Archs", age: 23},
		job:    "Software Engineer",
		salary: 75000.0,
	}

	// create a manager
	m := manager{
		employee: employee{
			person: person{name: "Russ Cox", age: 34},
			job:    "Lead Architect",
			salary: 104000.0,
		},
		teamSize: 10,
	}

	// print info
	fmt.Println(e.age)
	fmt.Println(e.job)
	fmt.Println(e.name)
	fmt.Println(e.person)
	fmt.Println(e.salary)

	fmt.Println(m.age)
	fmt.Println(m.job)
	fmt.Println(m.employee)
	fmt.Println(m.name)
	fmt.Println(m.teamSize)
	fmt.Println(m.salary)

}
