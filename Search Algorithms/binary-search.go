// Binary search is a more efficient search algorithm that works by repeatedly dividing the search interval in half. It requires the input array to be sorted beforehand.
package main

import "fmt"

func binarySeach(arr[]int, target int)int  {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println("Mitad del array", mid) // cue visual (no es parte del algoritmo)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target{
			low = mid + 1
		} else {
			high = mid -1
		}
	}
	return -1
}

func main() {
	arr := []int {1, 5, 7, 9, 12, 15}
	target := 44
	result := binarySeach(arr, target)

	if result == -1 {
		fmt.Println("Elemento no encontrado")
	} else {
		fmt.Printf("Elemento encontrado en index %d", result)
	}
	
}