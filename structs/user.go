package main

import "fmt"

type User struct {
	Name string
	Age int
	Email string
	Address string
}

func main() {
	user := User {
		Name: "John",
		Age: 30,
		Email: "john56@gmail.com",
		Address: "broklyn street 52",
	}
	fmt.Println(user)

}