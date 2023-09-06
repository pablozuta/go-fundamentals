// This program converts a temperature value to resistance using the Steinhart-Hart equation, simulating the behavior of a thermistor.

package main

import (
	"fmt"
	"math"
)

func temperatureToResistance(tempCelsius, referenceResistance, betaCoefficient float64) float64 {
	// Steinhart-Hart equation
	invTempKelvin := 1.0 / (tempCelsius + 273.15)
	return referenceResistance * math.Exp(betaCoefficient*(invTempKelvin-1))
}

func main() {
	fmt.Println("Temperature to resistance conversion (Thermistor)")

	var tempCelsius, referenceResistance, betaCoefficient float64
	fmt.Print("Enter temperature in celsius: ")
	fmt.Scanln(&tempCelsius)
	fmt.Print("Enter reference resistance (Ω): ")
	fmt.Scanln(&referenceResistance)
	fmt.Print("Enter beta coefficient: ")
	fmt.Scanln(&betaCoefficient)

	resistance := temperatureToResistance(tempCelsius, referenceResistance, betaCoefficient)
	fmt.Printf("Resistance at %.2fºC: %.2f Ω\n", tempCelsius, resistance)

}
