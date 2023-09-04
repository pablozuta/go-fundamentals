// This program calculates the voltage gain of a common-emitter BJT transistor amplifier.
package main

import "fmt"

func bjtAmplifierGain(vin, beta float64) float64 {
	return -beta * vin
}

func main() {
	fmt.Println("BTJ TRANSISTOR AMPLIFIER GAIN CALCULATOR")

	var vin, beta float64

	fmt.Print("Enter input voltage (V): ")
	fmt.Scanln(&vin)
	fmt.Print("Enter transistor current gain (β): ")
	fmt.Scanln(&beta)

	amplifierGain := bjtAmplifierGain(vin, beta)
	fmt.Printf("Amplifier Gain: %.2f\n", amplifierGain)

}
