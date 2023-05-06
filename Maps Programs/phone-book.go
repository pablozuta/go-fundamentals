package main

import "fmt"

func main()  {
	phoneBook := map[string]string {
		"Alice": "121-431-8323",
		"John": "312-122-3331",
		"Pharoah": "321-221-9831",
	}

	name := "John"
	phoneNumber, exists := phoneBook[name]

	if exists {
		fmt.Printf("%s's phone number is %s\n", name, phoneNumber)
	} else {
		fmt.Printf("No phone number found for %s\n", name) 
	}
}