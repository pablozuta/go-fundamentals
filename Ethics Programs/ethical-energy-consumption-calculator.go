// This program, the "Ethical Energy Consumption Calculator," assesses a user's monthly energy consumption and provides an ethical rating based on their energy efficiency. It encourages individuals to make ethical choices regarding energy conservation.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Ethical Energy Consumption Calculator!")
	fmt.Print("Enter your monthly energy consumption in kWh: ")

	var consumption float64
	fmt.Scan(&consumption)

	ethicalRating := assessEnergyConsumption(consumption)
	fmt.Printf("Ethical Rating For Energy Consumption: %s", ethicalRating)

}

func assessEnergyConsumption(consumption float64) string {
	if consumption < 0 {
		return "Invalid Input (Negative Consumption)"
	}

	if consumption < 200 {
		return "Low (efficient energy use)"
	} else if consumption < 500 {
		return "Moderate(consider reducing energy use)"
	} else {
		return "High (take inmediate action to reduce energy consumption)"
	}
}
