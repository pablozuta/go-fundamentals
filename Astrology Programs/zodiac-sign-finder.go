package main

import (
	"fmt"
	"time"
)

var zodiacSigns = []string{
	"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo", "Libra",
	"Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces",
}

var zodiacDates = []int{
	321, 420, 521, 621, 722, 823, 923,
	1023, 1122, 1222, 120, 219,
}

func findZodiacSign(month, day int) string {
	date := month*100 + day
	for i := len(zodiacDates) - 1; i >= 0; i-- {
		if date >= zodiacDates[i] {
			return zodiacSigns[i]
		}
	}
	return ""
}

func main() {
	currentTime := time.Now()
	month := int(currentTime.Month())
	day := currentTime.Day()
	zodiacSign := findZodiacSign(month, day)
	if zodiacSign != "" {
		fmt.Printf("Your zodiac sign is %s.\n", zodiacSign)
	} else {
		fmt.Println("Failed to determine your zodiac sign.")
	}
}
