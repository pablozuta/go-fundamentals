package main

import "fmt"

func selectionSort(arr[]int)  {
	for i := 0; i < len(arr) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func main()  {
	arr := []int {52, 54, 11, 84}
	fmt.Printf("Array original %v\n", arr)
	
	selectionSort(arr)
	fmt.Printf("Array sorted %v\n", arr)
}