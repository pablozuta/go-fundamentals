package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Title  string
	Author Author
}

type Author struct {
	FirstName string
	LastName  string
}
type ByLastName []Book

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].Author.LastName < a[j].Author.LastName }

func main() {
	books := []Book{
		{"The Great Gatsby", Author{"F. Scott", "Fitzgerald"}},
		{"Pride and Prejudice", Author{"Jane", "Austen"}},
		{"To Kill a Mockingbird", Author{"Harper", "Lee"}},
		{"1984", Author{"George", "Orwell"}},
		{"One Hundred Years of Solitude", Author{"Gabriel", "García Márquez"}},
	}
	// sort the books by author last name
	sort.Sort(ByLastName(books))

	// show the results
	for _, item := range books {
		fmt.Println(item)
	}

}
