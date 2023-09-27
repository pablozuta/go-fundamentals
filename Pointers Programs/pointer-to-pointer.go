// This program creates a variable x with an initial value of 10, creates a pointer ptr1 that points to x, creates a pointer ptr2 that points to ptr1, and then increments the value of x via the ptr2 pointer. Finally, it prints the updated value of x using both the x variable and ptr1.

package main

import "fmt"

func main() {
	x := 10
	ptr1 := &x
	ptr2 := &ptr1

	**ptr2++
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Value of x by a Pointer %d\n", *ptr1)
}
