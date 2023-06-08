package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int32 = 32
	p := unsafe.Pointer(&x)

	// manipulate pointer to add 1 to x
	pInt := (*int32)(p)
	*pInt += 1

	// print the value of x
	fmt.Println(x)
}