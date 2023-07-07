package main

import "fmt"

func IsSubset(setA, setB []int) bool {
	for _, elemA := range setA {
		isPresent := false
		for _, elemB := range setB {
			if elemA == elemB {
				isPresent = true
				break
			}
		}
		if !isPresent {
			return false
		}
	}
	return true
}

func main() {
	setA := []int{1, 2, 3}
	setB := []int{1, 2, 3, 4, 5}

	isSubset := IsSubset(setA, setB)

	fmt.Println("Set A:", setA)
	fmt.Println("Set B:", setB)
	fmt.Println("Is Set A a subset of Set B?", isSubset)

}
