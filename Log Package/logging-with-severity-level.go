package main

import (
	"log"
	"os"
)

func main() {
	// create a new logger with severity levels
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime)

	// Log messages with different severity levels
	logger.Println("Info: This is an informational message.")
	logger.Printf("Warning: %s\n", "This is a warning message.")
	logger.Fatalf("Error: %s\n", "This is a fatal error message.")
}
