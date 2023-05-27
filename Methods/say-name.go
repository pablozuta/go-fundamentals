package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) sayName() string {
	return "Hi, I'm " + p.Name
}

func main() {
	p := Person {Name: "Edgar Allan Poe"}
	fmt.Println(p.sayName())
}
