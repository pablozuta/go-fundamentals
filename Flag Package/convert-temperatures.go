// This program uses the Float64 function to define two command line flags: celsius and fahrenheit. It then uses flag.Parse() to parse the command line arguments.
// Depending on which flag was set, the program converts the temperature and prints the result.
package main

import (
	"flag"
	"fmt"
)

func main() {
	// define command-line flags (primer valor=flag , segundo valor=default value , tercero = descripcion)
	celsiusPtr := flag.Float64("celsius", 0.0, "temperature in celsius")
	fahrenheitPtr := flag.Float64("fahrenheit", 0.0, "temperature in fahrenheit")

	// parse command-line flags
	flag.Parse()

	// convert temperatures
	if *celsiusPtr != 0.0 {
		*fahrenheitPtr = (*celsiusPtr * 9.0 / 5.0) + 32.0
		fmt.Printf("The temperature %.2fº Celsius is %.2fº in Fahrenheit\n", *celsiusPtr, *fahrenheitPtr)
	} else if *fahrenheitPtr != 0 {
		*celsiusPtr = (*fahrenheitPtr - 32) * 5.0 / 9.0
		fmt.Printf("The temperature %.2fº Fahrenheit is %.2fº in Celsius\n", *fahrenheitPtr, *celsiusPtr)
	} else {
		fmt.Println("Please specify a temperature in Celsius or Fahrenheit")
	}
}

// to run the program
// $ go run main.go -celsius=25
// $ go run main.go -fahrenheit=19.4
