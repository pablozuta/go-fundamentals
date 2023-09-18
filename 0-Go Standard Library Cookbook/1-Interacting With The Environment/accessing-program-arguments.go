package main

import (
	"fmt"
	"os"
)

func main() {
	argumentos := os.Args

	// the first argument (zero index from slice) is the named of the binary(program)
	programName := argumentos[0]
	fmt.Printf("The name of the program is: %s\n", programName)

	// this call will print all command line arguments
	fmt.Println(argumentos)

	// the rest of the arguments can be obtained by omitting the first argument
	otherArgs := argumentos[1:]
	fmt.Println(otherArgs)

	// usar un for range para mostrar los argumentos y el index
	for index, item := range otherArgs {
		fmt.Printf("Index: %d Argument: %s \n", index, item)
	}
}
