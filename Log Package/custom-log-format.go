package main

import (
	"log"
	"os"
)

func main() {
	// create a new logger with a custom log format
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	// log a message
	logger.Println("Hello Logger!")
}
