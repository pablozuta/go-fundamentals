// This program restricts access to a resource based on the current time. It prompts the user to enter a start time and end time for the resource, and then checks the current time against the allowed time range. If the current time is within the allowed time range, the program grants access to the resource. Otherwise, it denies access.
package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Date(2023, time.April, 1, 9, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, time.April, 1, 17, 0, 0, 0, time.UTC)

	fmt.Println("Access Allowed from:", startTime.Format(time.RFC3339), "to", endTime.Format(time.RFC3339))

	for {
		currentTime := time.Now().UTC()
		if currentTime.After(startTime) && currentTime.Before(endTime) {
			fmt.Println("Access granted at", currentTime.Format(time.RFC3339))
			// Do something with the resource
			break
		} else {
			fmt.Println("Access denied at", currentTime.Format(time.RFC3339))
			time.Sleep(1 * time.Minute)
		}
	}
}
