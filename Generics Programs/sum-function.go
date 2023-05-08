package main

import "fmt"

func Sum[T any](slice []T, add func(T, T) T) T {
	var sum T
	for _, v := range slice {
		sum = add(sum, v)
	}
	return sum
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{2.4, 11.1, 421.2}

	sumInt := Sum(intSlice, func(x, y int) int { return x + y })
	sumFloat := Sum(floatSlice, func(x, y float64) float64 { return x + y })

	fmt.Println(sumInt)
	fmt.Println(sumFloat)
}
