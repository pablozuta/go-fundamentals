package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// declare and initialize arrays
	streetPrefixes := []string{
		"Neo",
		"Meta",
		"Cyber",
		"Tech",
		"Data",
		"Synth",
		"Hyper",
		"Tron",
		"Virt",
		"X",
	}
	streetSuffixes := []string{
		"Street",
		"Way",
		"Drive",
		"Avenue",
		"Lane",
		"Boulevard",
	}
	businnesTypes := []string{
		"Corporation",
		"Syndicate",
		"Agency",
		"Club",
		"Emporium",
		"Institute",
		"Industries",
		"Enterprises",
		"Firm",
		"Cooperative",
	}
	crimeSyndicates := []string{
		"Yakuza",
		"Triad",
		"Mafia",
		"Cartel",
		"Gang",
		"Sons of Adam",
		"Zaibatsu",
		"Tongs",
		"Vory",
		"Netrunners",
	}
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// generate a random street name
	streetName := streetPrefixes[rand.Intn(len(streetPrefixes))] + " " + streetSuffixes[rand.Intn(len(streetSuffixes))]
	// generate a syndicate
	businnesName := "Tokio Bar " + businnesTypes[rand.Intn(len(businnesTypes))] 
	// generate a mafia
	mafiaName := crimeSyndicates[rand.Intn(len(crimeSyndicates))]

	fmt.Printf("Welcome to %s\n", streetName)
	fmt.Printf("You are standing in front of %s\n", businnesName)
	fmt.Printf("Be careful, The %s crime syndicate is active in this area", mafiaName)
}
