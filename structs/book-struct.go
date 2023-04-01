package main

import "fmt"

type Book struct {
	title string
	author string
	numberOfPages int
}

func main() {
	book := Book {
		title: "Pedro Paramo",
		author: "Juan Rulfo",
		numberOfPages: 115,
	}
	fmt.Println("Title:", book.title)
	fmt.Println("Author:", book.author)
	fmt.Println("Number of pages:", book.numberOfPages)
}