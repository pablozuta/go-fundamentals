package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// The NewDense function is used to create a new 2x3 matrix a with the values {1, 2, 3, 4, 5, 6}. The []float64{1, 2, 3, 4, 5, 6} slice is passed as the data argument, and it represents the values of the matrix in row-major order.
	a := mat.NewDense(2, 3, [] float64 {1, 2, 3, 4, 5, 6})
	// The Gonum matrix library uses float64 as the default data type for matrix elements because float64 is a standard type for representing real numbers in scientific computing. 
	b := mat.NewDense(3, 2, [] float64 {7, 8, 9, 10, 11, 12})

	// multiplicar las matrices
	var result mat.Dense
	result.Mul(a, b)

	// mostrar los resultados
	fmt.Printf("  a * b   \n%v\n", mat.Formatted(&result))
}