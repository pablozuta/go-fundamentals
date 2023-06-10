package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	// Open input file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Create a mutex to synchronize access to wordCount map
	var mutex sync.Mutex

	// Create a map to store word frequencies
	wordCount := make(map[string]int)

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Create a buffered channel to limit concurrent goroutines
	concurrencyLimit := 5
	semaphore := make(chan struct{}, concurrencyLimit)

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		// Increment the wait group counter
		wg.Add(1)

		// Acquire a semaphore before launching a goroutine
		semaphore <- struct{}{}

		// Process each line concurrently
		go func(line string) {
			// Decrement the wait group counter when done
			defer func() {
				<-semaphore // Release the semaphore
				wg.Done()
			}()

			// Split the line into words
			words := strings.Fields(line)

			// Increment word frequencies
			for _, word := range words {
				lowercaseWord := strings.ToLower(word)

				// Acquire the mutex before accessing wordCount
				mutex.Lock()
				wordCount[lowercaseWord]++
				mutex.Unlock()
			}
		}(scanner.Text())
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo:", err)
		return
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the word frequencies
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

