// This program uses the `String` function to define a command line flag: `filename`. It then uses `flag.Parse()` to parse the command line arguments.
// The program opens the specified file, reads it line by line using a `bufio.Scanner`, and counts the number of lines, words, and characters in the file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// define command-line flags
	filenamePtr := flag.String("filename", "", "name of file to read")
	// parse command-line-flags
	flag.Parse()

	// open file
	file, err := os.Open(*filenamePtr)
	if err != nil {
		fmt.Println("Error Opening File", err)
		return
	}
	// diferir cierre de archivo
	defer file.Close()

	// read file, count lines, words and characters
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words, characters int
	for scanner.Scan() {
		words++
		characters += len(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Reading File", err)
		return
	}

	// Print results
	fmt.Printf("Words: %d\n", words)
	fmt.Printf("Characters: %d\n", characters)
}
