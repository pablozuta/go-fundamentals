// This program, "Ethical Donation Tracker," allows users to track donations to various recipients and promotes ethical giving practices.
package main

import "fmt"

type Donation struct {
	Recipient string
	Amount    float64
}

type DonationTracker struct {
	TotalDonated float64
	Donations    []Donation
}

func main() {
	tracker := DonationTracker{}
	fmt.Println("Welcome To The Ethical Donation Tracker!")

	for {
		fmt.Print("Enter recipient name(or 'exit' to finish): ")
		var recipient string
		fmt.Scan(&recipient)

		if recipient == "exit" {
			break
		}
		fmt.Print("Enter Donation Amount: ")
		var amount float64
		fmt.Scan(&amount)

		if amount <= 0 {
			fmt.Println("Invalid Donation Amount. Please enter a positive value.")
			continue
		}
		donation := Donation{Recipient: recipient, Amount: amount}
		tracker.AddDonation(donation)

	}
	fmt.Printf("Total Donated: $%.2f\n", tracker.TotalDonated)
	fmt.Println("Donation Details:")
	for _, donation := range tracker.Donations {
		fmt.Printf("%s: $%.2f\n", donation.Recipient, donation.Amount)
	}
}
func (tracker *DonationTracker) AddDonation(donation Donation) {
	tracker.TotalDonated += donation.Amount
	tracker.Donations = append(tracker.Donations, donation)
}
