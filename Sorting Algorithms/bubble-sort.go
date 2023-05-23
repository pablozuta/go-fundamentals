package main

import "fmt"

// bubbleSort performs bubble sort on the given integer slice.
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements and swap them if the current element is greater than the next element.
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap the elements using multiple assignment.
			}
		}
	}
}

func main() {
	arr := []int{52, 76, 32, 77, 1, 3, 55}

	// Call bubbleSort to sort the array in ascending order.
	bubbleSort(arr)

	fmt.Println("Sorted Array =", arr)
}
