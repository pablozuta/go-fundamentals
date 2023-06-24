package main

import (
	"fmt"
	"math/rand"
	"time"
)

type HorrorBook3 struct {
	Title  string
	Author string
	Year   int
	Rating float64
}

func main() {
	horrorBooks := []HorrorBook3{
		{Author: "Stephen King", Title: "It", Year: 1986, Rating: 4.5},
		{Author: "William Peter Blatty", Title: "The Exorcist", Year: 1971, Rating: 4.2},
		{Author: "Stephen King", Title: "The Shining", Year: 1977, Rating: 4.4},
		{Author: "Shirley Jackson", Title: "The Haunting of Hill House", Year: 1959, Rating: 4.1},
		{Author: "Josh Malerman", Title: "Bird Box", Year: 2014, Rating: 3.8},
		{Author: "Mark Z. Danielewski", Title: "House of Leaves", Year: 2000, Rating: 3.9},
		{Author: "Stephen King", Title: "The Stand", Year: 1978, Rating: 4.3},
		{Author: "Shirley Jackson", Title: "We Have Always Lived in the Castle", Year: 1962, Rating: 4.0},
		{Author: "Stephen King", Title: "The Outsider", Year: 2018, Rating: 4.1},
	}
	var favoriteAuthor string
	fmt.Printf("Enter the name of your favorite horror writer:")
	fmt.Scanln(&favoriteAuthor)

	var recommendedBook HorrorBook3
	for _, book := range horrorBooks {
		if book.Author == favoriteAuthor && book.Rating >= 4.0 {
			if recommendedBook.Title == "" {
				recommendedBook = book
			} else if rand.Float64() < 0.5 {
				recommendedBook = book
			}
		}
	}
	if recommendedBook.Title == "" {
		fmt.Printf("Sorry, we could not find any highly-rated horror books by %s.\n", favoriteAuthor)
	} else {
		fmt.Printf("Based on your favorite author, %s, we recommend the following book:\n", favoriteAuthor)
		fmt.Printf("- %s by %s (%d)\n", recommendedBook.Title, recommendedBook.Author, recommendedBook.Year)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
