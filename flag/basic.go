package main

import (
	"fmt"
	"flag"
)

func main()  {
	// definir command line flags
	wordPtr := flag.String("word", "foo", "a sting")
	numPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("bool", false, "a bool")

	// parsing
	flag.Parse()

	// acceder a los valores
	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("boolean:", *boolPtr)

	// podemos correr el programa asi:
	// go run basic.go -word=bar -num=123 -bool=true  
}