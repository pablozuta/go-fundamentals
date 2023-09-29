// In this program, we will create a struct PKDWork with five fields: Title, Year, Type, Rating, and Adaptations. We will create a slice of PKDWork structs and populate it with data on Philip K. Dick's books and movie adaptations. We will then prompt the user for a year range and a minimum rating and generate a list of works that match the user's criteria.
package main

import "fmt"

type PKDWork struct {
	Title       string
	Year        int
	Type        string
	Rating      float64
	Adaptations []string
}

func main() {
	pkdWorks := []PKDWork{
		{
			Title:       "Do Androids Dream of Electric Sheep?",
			Year:        1968,
			Type:        "book",
			Rating:      4.1,
			Adaptations: []string{"Blade Runner", "Blade Runner 2049"},
		},
		{
			Title:       "The Man in the High Castle",
			Year:        1962,
			Type:        "book",
			Rating:      4.0,
			Adaptations: []string{"The Man in the High Castle (TV series)"},
		},
		{
			Title:       "Ubik",
			Year:        1969,
			Type:        "book",
			Rating:      4.2,
			Adaptations: []string{},
		},
		{
			Title:       "A Scanner Darkly",
			Year:        2006,
			Type:        "movie",
			Rating:      4.0,
			Adaptations: []string{"A Scanner Darkly (book)"},
		},
		{
			Title:       "Blade Runner 2049",
			Year:        2017,
			Type:        "movie",
			Rating:      3.9,
			Adaptations: []string{},
		},
		{
			Title:       "The Adjustment Bureau",
			Year:        2011,
			Type:        "movie",
			Rating:      3.5,
			Adaptations: []string{"Adjustment Team (short story)"},
		},
	}

	var minRating float64
	fmt.Println("Enter a minumun rating (0.0 to 5.0)to filter works:")
	fmt.Scanln(&minRating)

	var startYear, endYear int
	fmt.Println("Enter a start year to filter works(e.g. 1960):")
	fmt.Scanln(&startYear)
	fmt.Println("Enter a end year to filter works(e.g. 1970):")
	fmt.Scanln(&endYear)

	var matchingTitles []string

	for _, work := range pkdWorks {
		if work.Rating >= minRating && work.Year >= startYear && work.Year <= endYear {
			matchingTitles = append(matchingTitles, work.Title)
		}
	}

	if len(matchingTitles) == 0 {
		fmt.Println("Sorry, no matching titles to show")
	} else {
		fmt.Printf("Matching Titles for ratings >= %.1f and years %d to %d:\n", minRating, startYear, endYear)

		for _, title := range matchingTitles {
			fmt.Println("-" + title + "year:")
		}
	}
}
