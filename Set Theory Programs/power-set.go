package main

import (
	"fmt"
	"math"
)

func PowerSet(set []int) [][]int {
	setSize := len(set)
	powerSetSize := int(math.Pow(2, float64(setSize)))

	powerSet := make([][]int, 0)

	for i := 0; i < powerSetSize; i++ {
		subset := make([]int, 0)
		for j := 0; j < setSize; j++ {
			if (i & (1 << uint(j))) != 0 {
				subset = append(subset, set[j])
			}
		}
		powerSet = append(powerSet, subset)
	}
	return powerSet
}

func main() {
	set := []int{1, 2, 3}

	powerSet := PowerSet(set)

	fmt.Println("Set:", set)
	fmt.Println("Power Set:", powerSet)
}
