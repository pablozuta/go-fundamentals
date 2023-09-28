// This program demonstrates how to use a pointer to a pointer to a struct in Go. It creates a Person struct with name and age fields, and creates a function updatePerson that takes a pointer to a pointer to a Person struct, and updates the age field of the Person struct to a new value. The function uses two levels of pointers so that it can modify the original pointer value (i.e., the pointer that was passed in to the function).

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func updatePerson(pp *person) {
	(*pp).age = 40
}

func main() {
	p := person{"Alice", 30}
	ptr1 := &p
	ptr2 := &ptr1

	updatePerson(**&ptr2)

	fmt.Printf("Name:%s - Age:%d", p.name, p.age)
}
