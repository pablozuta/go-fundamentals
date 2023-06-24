package main

import (
	"fmt"
	"sort"
)

// creamos un struct para usarlo en nuestro slice
type HorrorBook struct {
	Title           string
	Author          string
	PublicationYear int
}

func main() {
	// creamos un slice de structs HorrorBook
	horrorBooks := []HorrorBook{
		{Title: "It", Author: "Stephen King", PublicationYear: 1986},
		{Title: "The Exorcist", Author: "William Peter Blatty", PublicationYear: 1971},
		{Title: "The Shining", Author: "Stephen King", PublicationYear: 1977},
		{Title: "The Haunting of Hill House", Author: "Shirley Jackson", PublicationYear: 1959},
		{Title: "Bird Box", Author: "Josh Malerman", PublicationYear: 2014},
		{Title: "House of Leaves", Author: "Mark Z. Danielewski", PublicationYear: 2000},
	}
	// creamos un nuevo struct slice solo con los libros en los que el año es superior a 2000
	var recentHorrorBooks []HorrorBook
	for _, book := range horrorBooks {
		if book.PublicationYear > 2000 {
			// usamos el metodo append para ir agregando elementos al slice
			recentHorrorBooks = append(recentHorrorBooks, book)
		}
	}
	// hacemos sorting en ese nuevo slice usando el package sort ; ordenara los elementos de menor a mayor usando el field PublicationYear
	sort.Slice(recentHorrorBooks, func(i, j int) bool {
		return recentHorrorBooks[i].PublicationYear < recentHorrorBooks[j].PublicationYear
	})
	fmt.Println("Recent horror books:", recentHorrorBooks)

}
