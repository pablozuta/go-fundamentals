package main

import (
	"fmt"
	"strings"
)

const fraseRef = "Mery_has a little_lamb"

func main() {
	words := strings.Split(fraseRef, "_")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}

}
