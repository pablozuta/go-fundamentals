package main

import (
	"bufio"
	"fmt"
	"sort"
	"os"
	"strings"
	"sync"
)
type WordCount struct {
	Word string
	Count int
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("No se pudo abrir archivo", err)
		return
	}
	defer file.Close()

	// crear un mapa para guardar los resultados
	wordCount := make(map[string]int)

	// create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// iterate over each line of the file
	for scanner.Scan() {
		// increment the wait group counter
		wg.Add(1)

		// process each line concurrently
		go func (line string)  {
			// decrement the wait group counter when done
			defer wg.Done()

			// split the line into words
			words := strings.Fields(line)

			// increment word frequencies
			for _, word := range words {
				wordCount[strings.ToLower(word)]++
			}

		}(scanner.Text())
	}
	// check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al hacer scanner del archivo", err)
		return
	}

	// wait for all goroutines to finish
	wg.Wait()

	// filter word frequencies based on a threshold
	threshold := 3
	filteredWords := make([]WordCount, 0)
	for word, count := range wordCount {
		if count >= threshold {
			filteredWords = append(filteredWords, WordCount{word, count})
		}
	}
	// sort the filtered words in descending order
	sort.SliceStable(filteredWords, func(i, j int) bool {
		return filteredWords[i].Count > filteredWords[j].Count
	})

	// print the most frequent words
	fmt.Println("most frequent words:")
	for _, wc := range filteredWords {
		fmt.Printf("%s = %d times\n", wc.Word, wc.Count)
	}

}
