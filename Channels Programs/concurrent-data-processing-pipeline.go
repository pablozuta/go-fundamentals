package main

import "fmt"

// producer sends integers to a channel and then closes it
func producer(ch chan<- int) {
	// send integers from 0 to 9 to the channel
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// close the channel to signal that no more values will be sent
	close(ch)
}

// square receives integers from one channel, squares them, and sends them to another channel
func square(chIn <-chan int, chOut chan<- int) {
	// receive integers from the input channel until it is closed
	for v := range chIn {
		// square the integer and send the result to the output channel
		chOut <- v * v
	}
	// close the output channel to signal that no more values will be sent
	close(chOut)
}

// consumer receives integers from a channel and prints them to the console
func consumer(ch <-chan int) {
	// receive integers from the channel until it is closed
	for v := range ch {
		// print the integer to the console
		fmt.Println(v)
	}
}

func main() {
	// create two unbuffered channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// start a goroutine that sends integers to ch1
	go producer(ch1)

	// start a goroutine that squares integers received from ch1 and sends the results to ch2
	go square(ch1, ch2)

	// consume the squared integers from ch2 and print them to the console
	consumer(ch2)
}
