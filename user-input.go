// programa que lee input from user 
package main

import "fmt"

func main()  {
	fmt.Print("Enter a word:")
	var inputWord string

	// Scanln lee solo una palabra
	fmt.Scanln(&inputWord)

	fmt.Printf("The word was '%s'", inputWord)
}