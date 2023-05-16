package main

import (
	"fmt"
	"sort"
)

type Cyberware struct {
	name   string
	rarity string
}

func main() {
	// sample cyberware data
	cyberware := []Cyberware {
		{name: "Kerenzikov", rarity: "Legendary"},
		{name: "Mantis Blades", rarity: "Iconic"},
		{name: "Subdermal Grip", rarity: "Rare"},
		{name: "Gorilla Arms", rarity: "Uncommon"},
	}

	// sort cyberware by rarity (ordena alfabeticamente usando el string de rarity)
	sort.Slice(cyberware, func(i, j int) bool {
		return cyberware[i].rarity > cyberware[j].rarity
	})

	// print sorted cyberware
	fmt.Println("Sorted Cyberware:")
	for _, c := range cyberware {
		fmt.Printf("%s (%s)\n", c.name, c.rarity)
	}
}
