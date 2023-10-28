// This program, "Ethical Voting System Simulator," simulates a voting process with voters making ethical decisions when choosing a candidate. The program demonstrates ethical considerations in voting choices.
package main

import (
	"fmt"
)

type Voter struct {
	Name     string
	Ethical  bool
	Decision string
}

func main() {
	voters := []Voter{
		{Name: "Voter 1", Ethical: true},
		{Name: "Voter 2", Ethical: false},
		{Name: "Voter 3", Ethical: true},
	}

	simulateVoting(voters)
}

func simulateVoting(voters []Voter) {
	fmt.Println("Welcome to the Ethical Voting System Simulator!")

	for _, voter := range voters {

		if voter.Ethical {
			voter.Decision = "Supports an ethical candidate"
		} else {
			voter.Decision = "Do NOT support an ethical candidate"
		}
		fmt.Printf("%s, %s\n", voter.Name, voter.Decision)
	}

	fmt.Println("Simulation Complete...")

}
