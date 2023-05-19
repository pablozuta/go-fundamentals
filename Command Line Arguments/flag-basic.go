package main

import (
	"flag"
	"fmt"
)

func main() {
	// Declare variables for command-line flags
	var name string
	var age int

	// Define and parse command-line flags
	flag.StringVar(&name, "name", "World", "a name") // "World" valor default
	flag.IntVar(&age, "age", 0, "an age") // 0 valor default
	flag.Parse()

	// Print a personalized greeting
	fmt.Printf("Hello, %s! You are %d years old.", name, age)
}

// $ go run main.go -name Alice -age 30  
// go run main.go --help  (description of the flags)