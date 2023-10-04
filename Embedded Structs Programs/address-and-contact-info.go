// In this example, we have an Address struct and a ContactInfo struct that both embed an Email struct. The Email struct contains an email address, while the Address struct contains a street address, city, and state, and the ContactInfo struct contains a phone number.
package main

import "fmt"

type email struct {
	address string
}

type address struct {
	email  // embed the email struct
	street string
	city   string
	state  string
}

type contactInfo struct {
	email       // embed the email struct
	phoneNumber string
}

func main() {
	// create an address
	a := address{
		email:  email{address: "john.smith@gmail.com"},
		street: "123 Main Street",
		city:   "Anytown",
		state:  "CA",
	}

	// create a contact info
	c := contactInfo{
		email:       email{address: "jane.doe@hotmail.com"},
		phoneNumber: "555-1234",
	}

	// print some information about the address
	fmt.Println(a.address)
	fmt.Println(a.street)
	fmt.Println(a.city)
	fmt.Println(a.state)

	// print some information about the contactInfo
	fmt.Println(c.address)
	fmt.Println(c.phoneNumber)
}
