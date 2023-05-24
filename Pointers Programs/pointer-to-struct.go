package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	var p person
	var ptr *person

	ptr = &p
	ptr.name = "Alice"
	ptr.age = 34

	fmt.Printf("Name: %s, Age: %d\n", ptr.name, ptr.age)
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
	
}