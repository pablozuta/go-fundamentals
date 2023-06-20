package main

import "fmt"

func mergeSort(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}
	mid := n / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	numeros := []int{83, 22, 9, 77, 4, 11, 3, 1}
	numerosSorted := mergeSort(numeros)
	fmt.Println(numerosSorted)

}
