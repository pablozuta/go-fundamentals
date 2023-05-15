package main

import "fmt"

func producer(ch chan<- int)  {
	for i := 0; i < 10; i++ {
		fmt.Println("Sent from producer", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int)  {
	for v := range ch {
		fmt.Println("Show from consumer", v)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
	
}