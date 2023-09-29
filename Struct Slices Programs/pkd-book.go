// In this program, we will create a struct PKDBook with five fields: Title, Year, Rating, Adaptations, and Tags. We will create a slice of PKDBook structs and populate it with data on Philip K. Dick's books and their adaptations into movies and TV shows. We will then prompt the user for a tag and generate a list of book and adaptation titles that match the user's tag.

package main

import (
	"fmt"
)

type PKDBook struct {
	Title       string
	Year        int
	Rating      float64
	Adaptations []string
	Tags        []string
}

func main() {
	pkdBooks := []PKDBook{
		{
			Title:       "Do Androids Dream of Electric Sheep?",
			Year:        1968,
			Rating:      4.1,
			Adaptations: []string{"Blade Runner", "Blade Runner 2049"},
			Tags:        []string{"sci-fi", "dystopia", "androids"},
		},
		{
			Title:       "The Man in the High Castle",
			Year:        1962,
			Rating:      4.0,
			Adaptations: []string{"The Man in the High Castle (TV series)"},
			Tags:        []string{"alternate-history", "nazis", "japan"},
		},
		{
			Title:       "Ubik",
			Year:        1969,
			Rating:      4.2,
			Adaptations: []string{},
			Tags:        []string{"sci-fi", "time-travel", "death"},
		},
		{
			Title:       "A Scanner Darkly",
			Year:        1977,
			Rating:      4.0,
			Adaptations: []string{"A Scanner Darkly (film)"},
			Tags:        []string{"drug-culture", "police-state", "identity"},
		},
		{
			Title:       "Flow My Tears, the Policeman Said",
			Year:        1974,
			Rating:      4.0,
			Adaptations: []string{},
			Tags:        []string{"dystopia", "identity", "police-state"},
		},
	}

	var tag string
	fmt.Println("Enter a tag to find matching Philip K. Dick books and adaptations:")
	for _, item := range pkdBooks {
		for _, tag := range item.Tags {
			fmt.Println(tag)
		}
	}
	fmt.Println("")

	fmt.Scanln(&tag)

	var matchingTitles []string
	for _, book := range pkdBooks {
		if contains(book.Tags, tag) {
			matchingTitles = append(matchingTitles, book.Title)
			for _, adaptation := range book.Adaptations {
				matchingTitles = append(matchingTitles, adaptation)
			}
		}
	}

	if len(matchingTitles) == 0 {
		fmt.Printf("Sorry, not matches for tag \"%s\". \n", tag)
	} else {
		fmt.Printf("Matching titles for tag \"%s\": \n", tag)
		for _, title := range matchingTitles {
			fmt.Println("-" + title)
		}
	}

}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
