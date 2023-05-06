package main

import (
	"fmt"
	"math"
)

type Album struct {
	Title      string
	Artist     string
	Year       int
	TrackCount int
	Duration   float64 // in minutes
}

func main() {
	kindOfBlue := Album{
		Title:      "Kind of Blue",
		Artist:     "Miles Davis",
		Year:       1959,
		TrackCount: 5,
		Duration:   45.0,
	}
	fmt.Printf("Album: %s\n", kindOfBlue.Title)
	fmt.Printf("Artist: %s\n", kindOfBlue.Artist)
	fmt.Printf("Year: %d\n", kindOfBlue.Year)
	fmt.Printf("Track Count: %d\n", kindOfBlue.TrackCount)
	fmt.Printf("Duration: %.1f\n", kindOfBlue.Duration)

	// Calculate the average track length
    avgTrackLength := kindOfBlue.Duration / float64(kindOfBlue.TrackCount)
    fmt.Printf("Average track length: %.1f minutes\n", avgTrackLength)

    // Calculate the total playing time in seconds
    totalSeconds := math.Round(kindOfBlue.Duration * 60)
    fmt.Printf("Total playing time: %d seconds\n", int(totalSeconds))

}
