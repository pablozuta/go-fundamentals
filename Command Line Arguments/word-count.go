package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// open the file for reading
	file, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// create a scanner to read the file
	scanner := bufio.NewScanner(file)
	// set the scanner to split on words instead of lines
	scanner.Split(bufio.ScanWords)
	// count the number of words on file
	var count int
	for scanner.Scan() {
		count++
	}
	// mostrar el resultado
	fmt.Println("Numero de palabras en el archivo:", count)

}
