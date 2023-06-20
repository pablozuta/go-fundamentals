package main

import "fmt"

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	pivot := array[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, numero := range array[1:] {
		if numero <= pivot {
			left = append(left, numero)
		} else {
			right = append(right, numero)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func main() {
	numeros := []int{95, 54, 11, 4, 33, 1, 100, 31, 99}
	numerosSorted := quickSort(numeros)
	fmt.Println(numerosSorted)

}
