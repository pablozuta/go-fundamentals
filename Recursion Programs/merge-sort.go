package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	// means to slice the arr slice from the beginning (index 0) up to, but not including, the middle index. It includes all elements from index 0 to middle-1
	left := mergeSort(arr[:middle])
	// means to slice the arr slice starting from the middle index up to the end. It includes all elements from index middle to the last index of the slice.
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}

func main() {  
	arr := []int{5, 2, 4, 6, 1, 3}  
	fmt.Println("Unsorted Array:", arr)  
	  
	sortedArr := mergeSort(arr)  
	fmt.Println("Sorted Array:", sortedArr)  
}
