// This program is an "Environmental Carbon Footprint Calculator" that assesses a user's annual carbon emissions. It promotes ethical awareness by encouraging individuals to consider the environmental impact of their activities and take action to reduce their carbon footprint.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Environmental Carbon Footprint Calculator")
	fmt.Print("Please enter your annual carbon emissions in kilograms: ")

	var emissions float64
	fmt.Scan(&emissions)

	assessment := assessCarbonFootprint(emissions)
	fmt.Printf("Carbon Footprint Assessment: %s\n", assessment)
}

func assessCarbonFootprint(emissions float64) string {
	if emissions < 0 {
		return "Invalid Input (negative emissions)"
	}

	if emissions < 5000 {
		return "Low (environmental friendly)"
	} else if emissions < 10000 {
		return "Moderate (consider reducing)"
	} else {
		return "High (take inmediate action to reduce)"
	}

}
