package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Text:")
	scanner.Scan()
	text := scanner.Text()
	fmt.Printf("You entered: %v\n ", text)
}