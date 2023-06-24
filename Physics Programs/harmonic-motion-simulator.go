// This program simulates simple harmonic motion, which is a fundamental concept in physics. It uses the equation for harmonic motion x = A * sin(2πft + φ), where x is the displacement, A is the amplitude, f is the frequency, t is the time, and φ is the phase shift. The program calculates and prints the displacement at different time points over a specified duration.
package main

import (
	"fmt"
	"math"
)

func main() {
	amplitude := 1.0     // amplitude of the oscillation
	frequency := 2.0     // frequency of the oscillation
	phase := math.Pi / 2 // phase shift of the oscillation

	time := 0.0      // current time
	timeStep := 0.1  // time step size
	duration := 10.0 // total duration of the simulation

	for time <= duration {
		// calculate the displacement at the current time
		displacement := amplitude * math.Sin(2*math.Pi*frequency*time+phase)
		// print the time and displacement
		fmt.Printf("Time: %.2f, Displacement: %.2f\n", time, displacement)
		// increment time
		time += timeStep
	}
}

// This program can be used to visualize and understand the behavior of systems that exhibit simple harmonic motion, such as oscillating pendulums, vibrating strings, or mass-spring systems.
//It can be used as a teaching tool in physics courses to demonstrate the concept of simple harmonic motion and its mathematical representation.
