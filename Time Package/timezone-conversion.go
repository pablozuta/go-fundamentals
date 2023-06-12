package main

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2023-04-01T12:00:00.000Z"
	layout := "2006-01-02T15:04:05.000Z" 

	// parsing time in UTC
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("no se pudo hacer parsing", err)
		return
	}
	
	// convert time to Eastern Timezone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("no se pudo hacer parsing", err)
		return
	}

	easternTime := parsedTime.In(loc)
	
	fmt.Println("Parsed time in UTC: ", parsedTime)
	fmt.Println("Parsed time in Eastern Timezone: ", easternTime)

}