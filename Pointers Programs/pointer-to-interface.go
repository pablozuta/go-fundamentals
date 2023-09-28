// This program demonstrates how to use a pointer to an interface in Go. It creates a Shape interface with a getArea method, and two structs that implement the Shape interface: Circle and Rectangle. It then defines a getTotalArea function that takes a slice of Shape pointers, iterates over the slice, and calls the getArea method on each shape, summing up the areas of all shapes in the slice.

package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type circle struct {
	radius float64
}

type rectange struct {
	width, height float64
}

func (c *circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}
func (r *rectange) getArea() float64 {
	return r.width * r.height
}

func getTotalArea(shapes []shape) float64 {
	var totalArea float64

	for _, shape := range shapes {
		totalArea += shape.getArea()
	}
	return totalArea
}

func main() {
	shapes := []shape{
		&circle{radius: 5},
		&rectange{width: 10, height: 20},
	}
	fmt.Printf("Total Area %.2f\n", getTotalArea(shapes))
}
