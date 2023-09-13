package main

import (
	"fmt"
	"time"
)

type MoonPhase int

const (
	NewMoon MoonPhase = iota
	FirstQuarter
	FullMoon
	LastQuarter
)

func getMoonPhase(date time.Time) MoonPhase {
	year, month, day := date.Date()
	y := float64(year)
	m := float64(month)
	d := float64(day)
	c := 0.25

	// Calculation using the Conways algorithm
	epoch := float64(2451550.1 + 29.53058868*c)
	phase := (y*365.25 + m*30.6 + d - epoch) / 29.53058868
	phase = phase - float64(int(phase))
	phase *= 4

	switch {
	case phase < 0.5:
		return NewMoon
	case phase < 1.5:
		return FirstQuarter
	case phase < 2.5:
		return FullMoon
	default:
		return LastQuarter
	}
}

func main() {
	currentTime := time.Now()
	moonPhase := getMoonPhase(currentTime)

	switch moonPhase {
	case NewMoon:
		fmt.Println("The Moon is in the New Moon Phase")
	case FirstQuarter:
		fmt.Println("The Moon is in the First Quarter Phase")
	case FullMoon:
		fmt.Println("The Moon is in Full Moon phase")
	case LastQuarter:
		fmt.Println("The Moon is in the Last Quarter phase")
	default:
		fmt.Println("Failed to determine the moon phase")

	}
}
