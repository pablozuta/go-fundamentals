package main

import "fmt"

func main() {
	const (
		AgeOfUniverse = 13.8e9 // years
		SecondsInYear = 31557600
	)
	ageInSeconds := AgeOfUniverse * SecondsInYear
	fmt.Printf("The age of the universe is aproximately %.2e seconds", ageInSeconds)
}
