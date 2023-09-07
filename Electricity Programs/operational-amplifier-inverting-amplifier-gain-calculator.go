// This program calculates the voltage gain of an inverting amplifier circuit using an operational amplifier (op-amp).
package main

import "fmt"

func invertingAmplifierGain(r1, r2 float64) float64 {
	return -r2 / r1
}
func main() {
	fmt.Println("Operational Amplifier Inverting Amplifier Gain Calculator")

	var r1, r2 float64

	fmt.Print("Enter resistor R1 (Ω): ")
	fmt.Scanln(&r1)
	fmt.Print("Enter resistor R2 (Ω): ")
	fmt.Scanln(&r2)

	gain := invertingAmplifierGain(r1, r2)
	fmt.Printf("Voltage Gain: %.2f\n", gain)

}
