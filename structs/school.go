package main

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	school string
}

func main() {
	student := Student{Person: Person{name: "Bill Hicks", age: 30}, school: "Yale University"}
	fmt.Println("Name:", student.name)
	fmt.Println("Age:", student.age)
	fmt.Println("School:", student.school)
}