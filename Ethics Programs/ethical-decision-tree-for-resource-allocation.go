// This program is an "Ethical Decision Tree for Resource Allocation" that helps make decisions about allocating resources, such as a budget. It promotes ethical considerations in resource allocation by suggesting different allocation strategies based on the available budget.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Ethical Decision Tree for Resource Allocation!")
	fmt.Print("Enter the total budget available: ")

	var budget float64
	fmt.Scan(&budget)

	decision := makeResourceAllocationDecision(budget)

	fmt.Printf("Resource allocation decision: %s\n", decision)
}

func makeResourceAllocationDecision(budget float64) string {
	if budget < 0 {
		return "Invalid Input (negative budget)"
	}

	if budget < 1000 {
		return "Allocate resources primarily to basic necessities and social programs."
	} else if budget < 10000 {
		return "Balance funds between necessities, social programs and infrastructure."
	} else {
		return "Allocate funds evenly across necessities, social programs, infrastructure and sustainability initiatives."
	}
}
