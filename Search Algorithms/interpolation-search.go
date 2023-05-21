package main

import "fmt"

func interpolationSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for low <= high && x >= arr[low] && x <= arr[high] {
		pos := low + ((high-low)/(arr[high]-arr[low]))*(x-arr[low])

		if arr[pos] == x {
			return pos
		}
		if arr[pos] < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1

}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 12
	result := interpolationSearch(arr, x)
	if result == -1 {
		fmt.Printf("%d NO encontrado en array\n", x)
	} else {
		fmt.Printf("%d Encontrado en indice %d\n", x, result)
	}
}
