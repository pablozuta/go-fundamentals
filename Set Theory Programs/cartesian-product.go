package main

import "fmt"

type Pair struct {
	First  int
	Second int
}

func CartesianProduct(setA, setB []int) []Pair {
	cartesianProduct := make([]Pair, 0)

	for _, elemA := range setA {
		for _, elemB := range setB {
			pair := Pair{
				First:  elemA,
				Second: elemB,
			}
			cartesianProduct = append(cartesianProduct, pair)
		}
	}
	return cartesianProduct
}

func main() {
	setA := []int{1, 2}
	setB := []int{3, 4}

	product := CartesianProduct(setA, setB)

	fmt.Println("Set A:", setA)
	fmt.Println("Set B:", setB)
	fmt.Println("Cartesian Product:", product)

}
