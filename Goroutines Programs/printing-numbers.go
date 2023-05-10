package main

import (
	"fmt"
	"time"
)
func printNumbers(from, to int)  {
	for i := from; i <= to; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
	
}

func main() {
	go printNumbers(1, 5)
	go printNumbers(6, 10)
	time.Sleep(time.Second)
}