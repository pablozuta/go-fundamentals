// This program simulates an International Space Station (ISS) tracker. It repeatedly calls the `getISSLocation()` function to get simulated ISS coordinates (you would replace this with real data from an API or source). It then prints the current latitude and longitude of the ISS, updating every 10 seconds.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type location struct {
	latitude  float64
	longitude float64
}

func main() {
	iss := location{latitude: 51.5074, longitude: -0.1278} // London's Coordinates
	fmt.Printf("location of London %.2f\n", iss)

	for {
		issLocation := getISSLocation()
		fmt.Printf("ISS is currently at:\nLatitude: %.4f\nLongitude: %.4f\n\n", issLocation.latitude, issLocation.longitude)
		time.Sleep(10 * time.Second)
	}
}

func getISSLocation() location {
	// simulated function replace with real data
	return location{
		latitude:  28.6139 + (0.1*rand.Float64() - 0.5),
		longitude: -81.2090 + (0.1*rand.Float64() - 0.5),
	}
}
