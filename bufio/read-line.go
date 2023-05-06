package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Text: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Printf("You entered '%v'\n", text)

	scannerDos := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter more text: ")
	scannerDos.Scan()
	text = scannerDos.Text()
	fmt.Printf("Now you entered '%v'", text)
}