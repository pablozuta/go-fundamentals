package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{34, 21, 8, 90, 1, 2}
	insertionSort(arr)
	fmt.Println("Insertion Sort", arr)
}
