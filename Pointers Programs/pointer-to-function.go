// This program declares a function add that takes two integer parameters and returns their sum. It then declares a function apply that takes a function pointer f and two integer pointers x and y, and applies the function f to the values pointed to by x and y, updating the value pointed to by x with the result. Finally, it calls the apply function with the add function pointer and two integer pointers, and prints the updated value of the first integer pointer.
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func apply(f func(int, int) int, x, y *int) {
	result := f(*x, *y)
	*x = result
}

func main() {
	var a, b int = 10, 20
	ptr1 := &a
	ptr2 := &b

	apply(add, ptr1, ptr2)

	fmt.Printf("Value of a: %d\n", a)
}
