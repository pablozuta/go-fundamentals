// This program generates random numbers from a normal distribution using the Box-Muller transform. The NormalDistribution function takes the mean, standard deviation, and the number of samples as input, and returns a slice of random numbers following a normal distribution. The program then uses this function to generate and display a set of random samples from a normal distribution with a specified mean and standard deviation.
package main

import (
	"fmt"
	"math/rand"
)

func NormalDistribution(mean, stdDev float64, numSamples int) []float64 {
	samples := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		samples[i] = mean + stdDev*rand.NormFloat64()
	}
	return samples
}

func main() {
	mean := 0.0
	stdDev := 1.0
	numSamples := 10

	samples := NormalDistribution(mean, stdDev, numSamples)
	fmt.Println("Random samples from a normal distribution:")
	for _, item := range samples {
		fmt.Printf("%.6f\n", item)
	}
}

// Probability distributions are essential in statistical modeling and analysis. This program can be utilized whenever there is a need to generate random numbers from a normal distribution with specific mean and standard deviation. This can be useful in various scenarios such as simulating data for statistical experiments, generating random variables for testing algorithms, or generating synthetic datasets for machine learning tasks. The normal distribution is widely applicable in fields like finance, physics, and social sciences, where many phenomena follow a bell-shaped distribution.

// The program provides a convenient way to generate random samples from a normal distribution by specifying the mean, standard deviation, and the number of samples required. The generated samples can then be used for further analysis, such as calculating descriptive statistics, performing hypothesis tests, or visualizing the distribution using histograms or density plots.

// By adjusting the mean and standard deviation parameters, you can generate random samples that reflect different characteristics of the normal distribution, such as shifting the mean to model different central tendencies or changing the standard deviation to control the dispersion or spread of the data.
