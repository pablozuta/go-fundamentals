package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	numJobs := 10
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			worker(id, jobs, results)
			wg.Done()
		}(i)
	}

	// Send jobs to workers
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results from workers
	wg.Wait()
	close(results)

	// Print results
	for r := range results {
		fmt.Printf("Result: %d\n", r)
	}
}
