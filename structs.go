package main

import "fmt"

func main() {
	type Person struct {
		name string
		age int
	}

	person := Person{
		name: "Dave Garrison",
		age: 30,

	}
	fmt.Println(person.name)
	fmt.Println(person.age)

	}