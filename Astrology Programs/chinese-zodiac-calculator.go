package main

import (
	"fmt"
	"time"
)

var chineseZodiacSigns = []string{
	"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse",
	"Goat", "Monkey", "Rooster", "Dog", "Pig",
}

var chineseZodiacYears = []int{
	2020, 2021, 2022, 2023, 2024, 2025, 2026,
	2027, 2028, 2029, 2030, 2031,
}

func findChineseZodiacSign(year int) string {
	index := (year - 1900) % len(chineseZodiacSigns)
	if index < 0 {
		index += len(chineseZodiacSigns)
	}
	return chineseZodiacSigns[index]
}

func main() {
	currentTime := time.Now()
	year := currentTime.Year()
	chineseZodiacSign := findChineseZodiacSign(year)
	fmt.Printf("Your Chinese zodiac sign based on the year %d is %s.\n", year, chineseZodiacSign)

}
