package main

import (
	"fmt"
	
)

func main() {
	var x int = 42
	fmt.Printf("x= %d, address= %p\n ", x, &x)
}