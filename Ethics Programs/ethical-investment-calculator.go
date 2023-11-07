// This program, the "Ethical Investment Calculator," allows users to calculate the future value of their investments while emphasizing ethical considerations in investment decisions.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Ethical Investment Calculator!")
	fmt.Print("Enter your initial investment amount: $")

	var initialInvestment float64
	fmt.Scan(&initialInvestment)

	fmt.Print("Enter the annual return rate (as a decimal, e.g. 0.05 for 5%): ")
	var annualReturnRate float64
	fmt.Scan(&annualReturnRate)

	fmt.Print("Enter the investment duration(in years): ")
	var duration int
	fmt.Scan(&duration)

	finalAmount := calculateInvestment(initialInvestment, annualReturnRate, duration)
	fmt.Printf("After %d years, your investment will grow to: $%.2f\n", duration, finalAmount)
}

func calculateInvestment(initialInvestment, annualReturnRate float64, duration int) float64 {
	finalAmount := initialInvestment
	for i := 0; i < duration; i++ {
		finalAmount += finalAmount * annualReturnRate
	}
	return finalAmount
}
