package main

import "fmt"

type item struct {
	name  string
	price int
}

func main() {
	// sample item data
	inventory := []item{
		{name: "Kiroshi Optics Mk.3", price: 10_500},
		{name: "Satori", price: 3_500},
		{name: "Trajectory Analyser", price: 500},
	}

	totalPrice := 0
	for _, item := range inventory {
		totalPrice += item.price
	}
	fmt.Printf("Total Price of items in inventory : %d", totalPrice)
}
