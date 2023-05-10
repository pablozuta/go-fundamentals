package main

import "fmt"

func main() {
	data := "Neuromancer"
	bytes := []byte(data)
	fmt.Printf("%s\n", data)
	fmt.Printf("%v", bytes)
}