package main

import "fmt"

func main() {
	// declare a string
	str := "Hello, 世界!" 
	// print the string
	fmt.Println(str)
	// access individual characters
	fmt.Println(str[0]) // output 72
	fmt.Println(string(str[0])) // output H
	// iterate over the string
	for _, char := range str {
		fmt.Printf("%c ", char)
	}

}