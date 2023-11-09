// Go program inspired by William Gibson's "Count Zero" that incorporates the `slices`, `maps`, `structs`, `embedded structs`, `methods`, `pointers`, `interfaces`, `sorting`, and the `math` package for some calculations:

// We've added a `Score` field to the `Hacker` struct to represent the hacker's score.
// We've defined a sorting interface `ByScore` to sort hackers by their scores in descending order.
// We sort the hackers by their scores and print the top hackers.
// We calculate the average score of the hackers using the `math` package.
// We also perform a square root calculation using the `math` package.
package main

import (
	"fmt"
	"math"
	"sort"
)

// define an interface for a cyberdeck
type Cyberdeck2 interface {
	AccessMatrix2() string
}

// define a struct for a generic cyberdeck
type GenericCyberdeck2 struct {
	Model string
}

// implement the AccessMatrix method for the GenericCyberdeck
func (c *GenericCyberdeck2) AccessMatrix2() string {
	return "Accessing the matrix..."
}

// define a struct for a hacker character
type Hacker2 struct {
	Name       string
	Cyberdeck2 Cyberdeck2
	Score      int
}

// Create a method for the Hacker struct to access the Matrix
func (h *Hacker2) Hack() {
	fmt.Printf("%s is hacking with their %s\n", h.Name, h.Cyberdeck2.AccessMatrix2())
}

// define a sorting interface for hackers by their scores
type ByScore []*Hacker2

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].Score > s[j].Score }

func main() {
	// Create a map of hacker handles to their corresponding Hacker structs
	hackers := make(map[string]*Hacker2)

	// Create a slice of cyberdecks
	cyberdecks := []Cyberdeck2{
		&GenericCyberdeck2{Model: "Neuromancer 3000"},
		&GenericCyberdeck2{Model: "Decker's Delight"},
	}

	// Create hacker characters and assign them cyberdecks
	hacker1 := &Hacker2{Name: "Case", Cyberdeck2: cyberdecks[0], Score: 85}
	hacker2 := &Hacker2{Name: "Molly", Cyberdeck2: cyberdecks[1], Score: 92}

	// Add the hackers to the map
	hackers["Case"] = hacker1
	hackers["Molly"] = hacker2

	// Perform hacking operations
	for _, hacker := range hackers {
		hacker.Hack()
	}

	// Sort hackers by their scores in descending order
	hackerSlice := []*Hacker2{hacker1, hacker2}
	sort.Sort(ByScore(hackerSlice))

	fmt.Println("\nTop Hackers:")
	for i, hacker := range hackerSlice {
		fmt.Printf("%d. %s (Score: %d)\n", i+1, hacker.Name, hacker.Score)
	}

	// Calculate the average score of the hackers
	totalScore := 0
	for _, hacker := range hackers {
		totalScore += hacker.Score
	}
	averageScore := float64(totalScore) / float64(len(hackers))

	fmt.Printf("\nAverage Score: %.2f\n", averageScore)

	// Perform a math calculation (Square root)
	squareRoot := math.Sqrt(float64(totalScore))
	fmt.Printf("Square Root of Total Score: %.2f\n", squareRoot)

}
