package main

import "fmt"

func Intersection(setA, setB []int) []int {
	intersectionSet := make([]int, 0)

	// Check each element of setA if it exists in setB
	for _, elemA := range setA {
		for _, elemB := range setB {
			if elemA == elemB {
				intersectionSet = append(intersectionSet, elemA)
				break
			}
		}
	}
	return intersectionSet
}

func main() {
	setA := []int{1, 2, 3, 4, 5}
	setB := []int{4, 5, 6, 7, 8}

	intersectionSet := Intersection(setA, setB)

	fmt.Println("Set A:", setA)
	fmt.Println("Set B:", setB)
	fmt.Println("Intersection Set:", intersectionSet)

}
