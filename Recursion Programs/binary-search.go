// Binary search is a search algorithm that works by repeatedly dividing the search interval in half.
package main

import "fmt"

func binarySearch(arr []int, low int, high int, target int) int {
	// base case: if the search interval is empty return -1
	if high < low {
		return -1
	}

	// calculate the mid-point of the search interval
	mid := low + (high-low)/2

	// if the target is found at the mid point return the index
	if target == arr[mid] {
		return mid
	}
	// if the target is less than the mid point search the left half of the interval
	if arr[mid] > target {
		return binarySearch(arr, low, mid-1, target)
	}
	// if the target is greater than the mid point search the right half of the interval
	return binarySearch(arr, mid+1, high, target)

}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	target := 7
	index := binarySearch(arr, 0, len(arr)-1, target)
	fmt.Printf("The target value %d is found at index %d", target, index)
}
