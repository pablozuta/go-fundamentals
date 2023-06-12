package main

import "fmt"

// Binary search algorithm
func binarySearch(arr []int, low int, high int, x int) int {
	if high >= low {
		// Calculate the middle index
		mid := low + (high-low)/2

		if arr[mid] == x {
			// Element found at middle index
			return mid
		}

		if arr[mid] > x {
			// Search in the left half of the array
			return binarySearch(arr, low, mid-1, x)
		}

		// Search in the right half of the array
		return binarySearch(arr, mid+1, high, x)
	}

	// Element not found in the array
	return -1
}

// Exponential search algorithm
func exponentialSearch(arr []int, n int, x int) int {
	if arr[0] == x {
		// Element found at the beginning of the array
		return 0
	}

	i := 1
	// Double the index value (i) until it exceeds the array size or arr[i] becomes greater than x
	for i < n && arr[i] <= x {
		i *= 2
	}

	// Perform binary search within the identified range
	return binarySearch(arr, i/2, min(i, n-1), x)
}

// Utility function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example array and element to search
	arr := []int{2, 3, 4, 10, 40}
	x := 10
	n := len(arr)

	// Perform exponential search and store the result
	result := exponentialSearch(arr, n, x)

	if result == -1 {
		fmt.Println("Elemento no encontrado en array") // Element not found in the array
	} else {
		fmt.Printf("%d encontrado en array", x) // Element found in the array
	}
}
