package main

import (
	"fmt"
	"strings"
)

const refStringDos = "Mery had a little lamb"

func main() {
	words := strings.Fields(refStringDos)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}

}
