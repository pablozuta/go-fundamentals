package main

import (
	"fmt"
	"math"
)

const (
	numSteps2  = 1000   // Number of simulation steps
	timeStep2  = 0.01   // Time step size
	mass       = 1.0    // Mass of the particle
	omega      = 1.0    // Angular frequency
	initialPos = 1.0    // Initial position
	initialVel = 0.0    // Initial velocity
	hbar       = 1.0546 // Reduced Planck constant
)

func main() {
	// Initialize position and velocity
	position := initialPos
	velocity := initialVel

	// Perform quantum harmonic oscillator simulation
	for step := 0; step < numSteps2; step++ {
		// Calculate the wavefunction amplitude at the current position
		amplitude := math.Exp(-(mass * omega / (2 * hbar)) * position * position)

		// Update the position and velocity using the time evolution equations
		position = position + velocity*timeStep2
		velocity = velocity - (mass * omega * omega * position * timeStep2)

		// Normalize the wavefunction amplitude
		amplitude = amplitude / math.Sqrt(0.5*math.Pi)

		// Calculate the probability density at the current position
		probability := amplitude * amplitude

		// Print the current step and the probability density
		fmt.Printf("Step %d: Probability Density at Position %.2f = %.6f\n", step+1, position, probability)
	}

}
