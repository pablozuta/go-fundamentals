// This program defines two struct types: Author and Book. The Book struct embeds the Author struct using the syntax Author (with no field name). This means that the fields of the Author struct can be accessed directly from the Book struct.
package main

import "fmt"

type author struct {
	name string
	age  int
}

type book struct {
	title string
	author
	year int
}

func main() {
	// create a new book instance
	neuromancer := book{
		title: "Neuromancer",
		author: author{
			name: "William Gibson",
			age:  34,
		},
		year: 1984,
	}

	// print out the book information
	fmt.Printf("Title: %s\n", neuromancer.title)
	fmt.Printf("Name: %s\n", neuromancer.name)
	fmt.Printf("Year: %d\n", neuromancer.year)
}
