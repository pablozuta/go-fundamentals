package main

import (
	"fmt"
	"time"
)

func main() {
	var durationStr string
	fmt.Print("Enter a duration (e.g. 4s, 2m40s): ")
	fmt.Scanln(&durationStr)

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("Unable to read line", durationStr)
		return
	}

	fmt.Println("Countdown started for:", duration)

	startTime := time.Now()
	endTime := startTime.Add(duration)
	remaining := duration

	for remaining > 0 {
		time.Sleep(time.Second)
		remaining = endTime.Sub(time.Now())
		fmt.Println("Time remaining:", remaining)
	}

	fmt.Println("Time's Up")



}