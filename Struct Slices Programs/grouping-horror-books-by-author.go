package main

import "fmt"

type HorrorBook2 struct {
	Title           string
	Author          string
	PublicationYear int
}

func main() {
	horrorBooks := []HorrorBook2{
		{Title: "It", Author: "Stephen King", PublicationYear: 1986},
		{Title: "The Exorcist", Author: "William Peter Blatty", PublicationYear: 1971},
		{Title: "The Shining", Author: "Stephen King", PublicationYear: 1977},
		{Title: "The Haunting of Hill House", Author: "Shirley Jackson", PublicationYear: 1959},
		{Title: "Bird Box", Author: "Josh Malerman", PublicationYear: 2014},
		{Title: "House of Leaves", Author: "Mark Z. Danielewski", PublicationYear: 2000},
	}

	booksByAuthor := make(map[string][]string)
	for _, book := range horrorBooks {
		booksByAuthor[book.Author] = append(booksByAuthor[book.Author], book.Title)
	}
	for author, titles := range booksByAuthor {
		fmt.Println(author)
		for _, title := range titles {
			fmt.Println("-", title)
		}
	}
}
