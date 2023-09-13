package main

import (
	"fmt"
	"time"
)

var sunSigns = []string{
	"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo", "Libra",
	"Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces",
}

var sunSignStarDates = []time.Time{
	time.Date(0, time.March, 21, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.April, 20, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.May, 21, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.June, 21, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.July, 23, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.August, 23, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.September, 23, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.October, 23, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.November, 22, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.December, 22, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.January, 20, 0, 0, 0, 0, time.UTC),
	time.Date(0, time.February, 19, 0, 0, 0, 0, time.UTC),
}

func findSunSign(date time.Time) string {
	year := 0 // leap year for February 29th check
	for i := len(sunSignStarDates) - 1; i >= 0; i-- {
		if date.After(sunSignStarDates[i]) || date.Equal(sunSignStarDates[i]) {
			if i == 11 && (date.Month() == time.March && date.Day() == 20 || date.Month() == time.March && date.Day() == 21 && date.Hour() < 12) {
				year-- // special case : February 29th
			}
			return sunSigns[(date.Year()+year)%len(sunSigns)]
		}
	}

	return ""

}

func main() {
	currentTime := time.Now()
	sunSign := findSunSign(currentTime)
	if sunSign != "" {
		fmt.Printf("Your sun sign is %s", sunSign)
	} else {
		fmt.Println("Failed to determine your sun sign")
	}
}
