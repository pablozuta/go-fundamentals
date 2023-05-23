package main

import "fmt"

type Record struct {
	title     string
	author    string
	year      int
	trackList []string
	personel  []string
}

func main() {
	aLoveSupreme := Record{
		title:     "A Love Supreme",
		author:    "John Coltrane",
		year:      1964,
		trackList: []string{"Acknoledgement", "Resolution", "Pursuance", "Psalm"},
		personel:  []string{"John Coltrane", "McCoy Tyner", "Jimmy Garrison", "Elvin Jones"},
	}
	fmt.Println("Title - ", aLoveSupreme.title)
	fmt.Println("Author - ", aLoveSupreme.author)
	fmt.Println("Year - ", aLoveSupreme.year)
	for _, track := range aLoveSupreme.trackList {

		fmt.Println("Track - ", track)
	}
	fmt.Println("Personel - ", aLoveSupreme.personel)

}
