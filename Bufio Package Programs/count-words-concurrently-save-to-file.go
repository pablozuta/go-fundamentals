package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

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
			defer wg.Done()

			// Split the line into words
			words := strings.Fields(line)

			// Increment word frequencies
			for _, word := range words {
				wordCount[strings.ToLower(word)]++
			}

			// Release the semaphore
			<-semaphore
		}(scanner.Text())
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Filter word frequencies based on a threshold
	threshold := 3
	filteredWords := make([]WordCount, 0)
	for word, count := range wordCount {
		if count >= threshold {
			filteredWords = append(filteredWords, WordCount{word, count})
		}
	}

	// Sort the filtered words by count in descending order
	sort.Slice(filteredWords, func(i, j int) bool {
		return filteredWords[i].Count > filteredWords[j].Count
	})

	// Create a new file to save the most frequent words
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	// Create a writer to write to the output file
	writer := bufio.NewWriter(outputFile)

	// Write the most frequent words to the output file
	for _, wc := range filteredWords {
		line := fmt.Sprintf("%s: %d\n", wc.Word, wc.Count)
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Flush the writer to ensure all data is written
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	fmt.Println("Most frequent words saved to output.txt.")
}