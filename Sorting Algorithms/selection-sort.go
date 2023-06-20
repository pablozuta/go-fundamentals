package main

import "fmt"

func selectionSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min_idx] {
				min_idx = j
			}
		}
		array[i], array[min_idx] = array[min_idx], array[i]
	}

}

func main() {
	numeros := []int{89, 43, 73, 1, 44, 11, 3}
	selectionSort(numeros)
	fmt.Println(numeros)
}
