// This program demonstrates linear regression, a statistical technique used to model the relationship between two variables by fitting a linear equation to observed data points. The LinearRegression function takes two slices of data points (x and y) and calculates the slope and intercept of the best-fit line. The program then uses these values to perform linear regression on a provided dataset and displays the slope and intercept of the resulting line.
package main

import (
	"fmt"
)

func LinearRegression(x, y []float64) (float64, float64) {
	n := len(x)
	sumX, sumY, sumXY, sumXX := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXX += x[i] * x[i]
	}
	// calculate the slope(m)
	slope := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXX - sumX*sumX)

	// calculate the intercept(c)
	intercept := (sumY - slope*sumX) / float64(n)

	return slope, intercept

}

func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	slope, intercept := LinearRegression(x, y)

	fmt.Printf("Slope (m): %.2f\n", slope)
	fmt.Printf("Intercept (c): %.2f\n", intercept)

}

// Linear regression is widely used in various fields, such as economics, finance, and data analysis, to understand the relationship between variables and make predictions. This program can be utilized whenever there is a need to perform linear regression on a set of data points and determine the equation of the best-fit line. For example, it could be used to analyze the relationship between advertising spending and sales revenue or to predict housing prices based on features like area and number of bedrooms.
