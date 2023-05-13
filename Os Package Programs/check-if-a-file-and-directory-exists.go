package main

import (
	"fmt"
	"os"
)

func main() {
	// check if a file exist
	fileName := "archivo.txt"
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("%s exist!\n", fileName)
	} else {
		fmt.Printf("%s DOES NOT exist!\n", fileName)

	}

	// check if a directory exist
	dirName := "test"
	if _, err := os.Stat(dirName); err == nil {
		fmt.Println("Directory exist.")
	} else {
		fmt.Println("Directory DOES NOT exist.")

	}
}
