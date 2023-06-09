package main

import "fmt"

func main() {
	// byte slice
	byteSlice := []byte {70, 109, 60, 20}

	// string
	str := "Cybernetics"

	// print values
	fmt.Println(byteSlice)
	fmt.Println(str)
	
	// modifying the byte slice
	byteSlice[0] = 23
	
	// modifying the string
	str = "Neo" + str[0:]

	// print values
	fmt.Println(byteSlice)
	fmt.Println(str)
}