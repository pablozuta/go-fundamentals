// This program uses the Monte Carlo method to estimate the value of Pi. The Monte Carlo method is a statistical technique that uses random sampling to solve problems. In this case, it randomly generates points within a square and checks if they fall inside a unit circle centered at the origin. By comparing the number of points inside the circle to the total number of points generated, an approximation of Pi can be obtained.
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	totalPoints := 1000000
	pointsInsideCircle := 0

	for i := 0; i < totalPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()

		distance := math.Sqrt(x*x + y*y)
		if distance <= 1 {
			pointsInsideCircle++
		}
	}
	piEstimation := 4 * float64(pointsInsideCircle) / float64(totalPoints)
	fmt.Printf("Estimated value of Pi: %.6f\n", piEstimation)
}

// Pi estimation: The program can be used to estimate the value of Pi. The more points generated, the more accurate the estimation becomes. This technique is particularly useful when mathematical or analytical approaches to calculating Pi are challenging or impractical.  
// Statistical simulations: The Monte Carlo method has applications in various fields, such as finance, engineering, and physics. It can be used to simulate and analyze complex systems, model uncertainties, and solve problems that involve probabilistic calculations.  
// History:  
// The Monte Carlo method was named after the famous casino in Monaco and was first developed by scientists working on the Manhattan Project during World War II. It was initially used to solve problems related to the development of atomic bombs. Since then, the Monte Carlo method has found widespread use in various scientific and engineering disciplines for solving complex problems and making estimations.  
  
// The use of Monte Carlo simulations, such as estimating Pi using random sampling, has become a common technique to illustrate the concept and application of the Monte Carlo method. These simulations provide a practical and intuitive way to understand the probabilistic nature of the method and its ability to approximate solutions. 
