package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime)

	if true {
		logger.Println("This message will be logged")
	}
	if false {
		logger.Println("This message won't be logged")
	}
}
