package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration, ch chan<- bool)  {
	time.Sleep(duration)
	ch <- true
}
func main() {
	ch := make(chan bool)
	duration := 5 * time.Second
	fmt.Printf("Starting timer for %v\n", duration)
	go timer(duration, ch)
	<-ch
	fmt.Println("Time Complete")

}