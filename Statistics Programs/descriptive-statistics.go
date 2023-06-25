// This program demonstrates the calculation of mean and standard deviation, which are fundamental concepts in descriptive statistics. The CalcMean function calculates the average value of a given slice of numbers, while the CalcStandardDeviation function computes the standard deviation, which measures the dispersion or spread of the data points around the mean. The program then uses these functions to calculate and display the mean and standard deviation of a provided dataset.
package main

import (
	"fmt"
	"math"
)

// calculates the mean of a given slice of numbers
func calcMean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// calculates the standard deviation of a given slice of numbers
func calcStandardDeviation(data []float64) float64 {
	mean := calcMean(data)
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-mean, 2)
	}
	variance := sum / float64(len(data))
	return math.Sqrt(variance)
}

func main() {
	data := []float64{4, 5, 2, 8, 9, 3, 6, 1, 7}
	mean := calcMean(data)
	standardDeviation := calcStandardDeviation(data)

	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Standard Deviation: %.2f\n", standardDeviation)
}

// Descriptive statistics is widely used in various fields, such as finance, economics, social sciences, and data analysis. This program can be utilized whenever mean and standard deviation need to be calculated for a set of data points. For example, it could be used to analyze the average sales of a product over a period of time or to measure the variability of test scores in a classroom.
