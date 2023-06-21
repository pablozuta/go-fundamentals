package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// create a new error with stack trace
	err := errors.New("Algo Salio Mal")
	log.Println(fmt.Errorf("%w", err))
}
