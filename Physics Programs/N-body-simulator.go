package main

import (
	"fmt"
	"math"
)

const (
	numParticles = 2 // Number of particles
	numSteps     = 1000
	timeStep     = 0.01      // Time step size
	gravConstant = 6.674e-11 // Gravitational constant (m^3 kg^−1 s^−2)
)

type Particle struct {
	x, y, mass, vx, vy float64 // Position, mass, and velocity components
}

func main() {
	particles := make([]Particle, numParticles)

	// Initialize particle properties
	particles[0] = Particle{x: 0, y: 0, mass: 5.972e24, vx: 0, vy: 0}          // Earth
	particles[1] = Particle{x: 384.4e6, y: 0, mass: 7.348e22, vx: 0, vy: 1022} // Moon

	// Perform N-body simulation
	for step := 0; step < numSteps; step++ {
		// Update particle positions and velocities
		for i := 0; i < numParticles; i++ {
			for j := 0; j < numParticles; j++ {
				if i != j {
					dx := particles[j].x - particles[i].x
					dy := particles[j].y - particles[i].y
					r := math.Sqrt(dx*dx + dy*dy)
					force := gravConstant * particles[i].mass * particles[j].mass / (r * r)
					theta := math.Atan2(dy, dx)
					fx := force * math.Cos(theta)
					fy := force * math.Sin(theta)

					particles[i].vx += fx / particles[i].mass * timeStep
					particles[i].vy += fy / particles[i].mass * timeStep
				}
			}
		}

		for i := 0; i < numParticles; i++ {
			particles[i].x += particles[i].vx * timeStep
			particles[i].y += particles[i].vy * timeStep
		}
	}

	// Print final particle positions
	for i := 0; i < numParticles; i++ {
		fmt.Printf("Particle %d: Position (%.2f, %.2f)\n", i, particles[i].x, particles[i].y)
	}

}
