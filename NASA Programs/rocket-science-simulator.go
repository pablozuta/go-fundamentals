// This program simulates a basic rocket launch and descent. It calculates and prints the rocket's position over time using the kinematic equation for motion. The rocket starts from the ground with an initial velocity of 0 m/s and accelerates upward with a constant acceleration of 10 m/s² until it reaches its peak height and then descends back to the ground. The simulation stops when the rocket lands.

package main

import (
	"fmt"
	"math"
)

func main() {
	initialVelocity := 0.0
	acceleration := 10.0

	time := 0.0
	position := 0.0

	fmt.Println("Rocket Science Simulator")

	for position >= 0 {
		position = calculatePosition(initialVelocity, acceleration, time)
		fmt.Printf("Time: %.2f s |Position: %2.f m\n ", time, position)
		time += 0.1
	}
}

func calculatePosition(initialVelocity, acceleration, time float64) float64 {
	return initialVelocity*time + 0.5*acceleration*math.Pow(time, 2)
}
