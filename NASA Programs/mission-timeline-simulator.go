// This program simulates a timeline of events for a space mission. It defines a series of `MissionEvent` structs, each with a name and a duration. The program then iterates through the events, printing the time and event name, waiting for the specified duration, and updating the current time.

package main

import (
	"fmt"
	"time"
)

type MissionEvent struct {
	Name     string
	Duration time.Duration
}

func main() {
	missionEvents := []MissionEvent{
		{"Launch", 10 * time.Second},
		{"Orbit Insertion", 20 * time.Second},
		{"Scientific Observations", 3 * time.Minute},
		{"Data Transmission", 30 * time.Second},
		{"Reentry", 15 * time.Second},
		{"Mission Complete", 5 * time.Second},
	}

	currentTime := time.Now()
	fmt.Println("Mission Timeline Program")
	fmt.Println("")

	for _, event := range missionEvents {
		fmt.Printf("%s - %s\n", currentTime.Format("15:04:05"), event.Name)
		time.Sleep(event.Duration)
		currentTime = currentTime.Add(event.Duration)
	}

	fmt.Println("Mission Finished")
}
