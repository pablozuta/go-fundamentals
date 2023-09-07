// This program calculates the frequency of an astable multivibrator circuit using a 555 timer IC.

package main

import (
	"fmt"
)

func astableFrequency(r1, r2, c float64) float64 {
	return 1.44 / ((r1 + 2*r2) * c)
}

func main() {
	fmt.Println("555 Timer Astable Oscillator Frequency Calculator")

	var r1, r2, c float64

	fmt.Print("Enter Resistor R1 (Ω): ")
	fmt.Scanln(&r1)
	fmt.Print("Enter Resistor R2 (Ω): ")
	fmt.Scanln(&r2)
	fmt.Print("Enter Capacitor C (F):")
	fmt.Scanln(&c)

	frequency := astableFrequency(r1, r2, c)
	fmt.Printf("Frequency: %.2f Hz\n", frequency)

}
