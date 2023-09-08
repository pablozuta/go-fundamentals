package main

import (
	"fmt"
	"time"
)

func main() {
	launchTime := time.Date(2023, 7, 10, 15, 30, 0, 0, time.UTC)
	currentTime := time.Now()

	fmt.Println("Countdown Space Launch")

	for currentTime.Before(launchTime) {
		timeLeft := launchTime.Sub(currentTime)
		fmt.Printf("%02d:%02d:%02d\n", int(timeLeft.Hours()), int(timeLeft.Minutes())%60, int(timeLeft.Seconds())%60)
		time.Sleep(1 * time.Second)
		currentTime = time.Now()
	}
	fmt.Println("Liftoff!!")
}
