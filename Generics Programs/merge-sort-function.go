package main

import "fmt"

func mergeSort[T comparable](slice []T, less func(a, b T) bool) []T {
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSort(slice[:mid], less)
	right := mergeSort(slice[mid:], less)

	return merge(left, right, less)
}

func merge[T comparable](left, right []T, less func(a, b T) bool) []T {
	result := make([]T, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			result = append(result, right...)
		} else if len(right) == 0 {
			result = append(result, left...)
		} else if less(left[0], right[0]) {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func main() {
	intSlice := []int{5, 2, 7, 1, 8}
	lessInt := func(a, b int) bool { return a < b }
	sortedInts := mergeSort(intSlice, lessInt)
	fmt.Println(sortedInts)
	stringSlice := []string{"foo", "bar", "baz", "qux"}
	lessString := func(a, b string) bool { return a < b }
	sortedStrings := mergeSort(stringSlice, lessString)
	fmt.Println(sortedStrings) // Output: [bar baz foo qux]
}


