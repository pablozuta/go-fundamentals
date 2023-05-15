package main

import (
	"fmt"
	"time"
)

// This function sends numbers from 0 to 9 to the channel `ch`,
// one number per second. After sending all the numbers, it closes the channel.
func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i // Send `i` to the channel `ch`
		time.Sleep(time.Second) // Wait for one second
	}
	close(ch) // Close the channel after all the numbers are sent
}

// This function receives numbers from the channel `ch` and prints them.
// It keeps receiving numbers until the channel is closed.
func consumer(ch <-chan int) {
	for {
		// Check if the channel is open or closed and receive a value from it.
		// If the channel is closed, the second return value `ok` will be `false`.
		if v, ok := <-ch; ok {
			fmt.Println("Received", v) // Print the received value
		} else {
			fmt.Print("Channel is Closed and Stoped sending Numbers")
			break // Exit the loop when the channel is closed
		}
	}
}

func main() {
	ch := make(chan int) // Create an unbuffered channel of integers
	go producer(ch) // Start the producer goroutine
	go consumer(ch) // Start the consumer goroutine
	time.Sleep(11 * time.Second) // Wait for 11 seconds for the goroutines to finish
}
