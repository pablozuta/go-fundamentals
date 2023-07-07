package main

import "fmt"

func Union(setA, setB []int) []int {
	unionSet := make([]int, 0)

	// add elements from setA to unionSet
	for _, element := range setA {
		unionSet = append(unionSet, element)
	}
	// Add elements from setB to unionSet, if they are not already present
	for _, element := range setB {
		isDuplicate := false

		for _, existingElement := range unionSet {
			if element == existingElement {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			unionSet = append(unionSet, element)
		}
	}
	return unionSet
}
func main() {
	setA := []int{1, 2, 3, 4, 5}
	setB := []int{4, 5, 6, 7, 8}

	unionSet := Union(setA, setB)

	fmt.Println("Set A:", setA)
	fmt.Println("Set B:", setB)
	fmt.Println("Union Set:", unionSet)
}
