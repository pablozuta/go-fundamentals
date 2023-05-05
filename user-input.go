// programa que lee input from user 
package main

import "fmt"

func main()  {
	fmt.Print("Enter a phrase:")
	var inputPhrase string
	// Scanln lee solo una palabra
	fmt.Scanln(&inputPhrase)

	fmt.Printf("Phrase was '%s'", inputPhrase)
}