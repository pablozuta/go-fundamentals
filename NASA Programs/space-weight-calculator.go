package main

import "fmt"

func main() {
	fmt.Print("Enter your weight (kg): ")
	var weight float64
	_, err := fmt.Scan(&weight)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	earthGravity := 9.81
	marsGravity := 3.72

	earthWeight := weight * earthGravity
	marsWeight := weight * marsGravity

	fmt.Printf("Your weight on Earth: %.2f N\n", earthWeight)
	fmt.Printf("Your weight on Mars: %.2f N\n", marsWeight)

}
