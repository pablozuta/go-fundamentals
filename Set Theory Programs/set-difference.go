package main

import "fmt"

func Difference(setA, setB []int) []int {
	differenceSet := make([]int, 0)

	// Check each element of setA if it does not exist in setB
	for _, elemA := range setA {
		isPresent := false
		for _, elemB := range setB {
			if elemA == elemB {
				isPresent = true
				break
			}
		}
		if !isPresent {
			differenceSet = append(differenceSet, elemA)
		}
	}
	return differenceSet

}

func main() {
	setA := []int{1, 2, 3, 4, 5}
	setB := []int{4, 5, 6, 7, 8}

	differenceSet := Difference(setA, setB)

	fmt.Println("Set A:", setA)
	fmt.Println("Set B:", setB)
	fmt.Println("Set Difference:", differenceSet)

}
