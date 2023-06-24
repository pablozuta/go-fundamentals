// This program simulates projectile motion, which describes the motion of objects thrown into the air under the influence of gravity. It calculates and prints the position of a projectile at different time points using the equations of motion.
package main

import (
	"fmt"
	"math"
)

const gravity = 9.8 // acceleration due to gravity

type Projectile struct {
	x, y, time float64 // position coordinates and time
	vx, vy     float64 // velocity component in x and y directions
}

func main() {
	initialVelocity := 25.0 // initial velocity for the projectile (m/s)
	angle := math.Pi / 4    // launche angle of the projectile (radians)
	timeStep := 0.1         // The time step for each iteration (s)

	projectile := Projectile{
		x:    0,
		y:    0,
		vx:   initialVelocity * math.Cos(angle),
		vy:   initialVelocity * math.Sin(angle),
		time: 0,
	}
	for projectile.y >= 0 {
		// update the projectile's position and velocity
		projectile.x += projectile.vx * timeStep
		projectile.y += projectile.vy * timeStep
		projectile.vy -= gravity * timeStep

		// Print the current position
		fmt.Printf("Time: %.2f, Position: (%.2f, %.2f)\n", projectile.time, projectile.x, projectile.y)

		projectile.time += timeStep // Increment the time
	}

}
// 
