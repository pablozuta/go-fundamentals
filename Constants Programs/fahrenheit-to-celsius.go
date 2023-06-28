// In this program, we define a constant freezingFahrenheit with a value of 32.0 (the freezing point of water in Fahrenheit). We then convert a temperature in Fahrenheit (68.0) to Celsius using the formula (fahrenheit - freezingFahrenheit) * 5 / 9.
package main

import "fmt"

func main() {
	const freezingFahrenheit = 32.0

	fahrenheit := 68.0
	celsius := (fahrenheit - freezingFahrenheit) * 5 / 9
	fmt.Printf("%.2fºF is equal to %.2fºC\n", fahrenheit, celsius)

}
