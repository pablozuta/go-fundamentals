package main

import "fmt"

type Quest struct {
	name     string
	npcs     []NPC
	duration int
}

type NPC struct {
	name  string
	level int
}

func main() {
	// sample quest data
	quests := []Quest{
		{name: "The Heist",
			duration: 60,
			npcs: []NPC{
				{name: "Dexter Deshawn", level: 10},
				{name: "Evelyn Parker", level: 12},
				{name: "Jackie Welles", level: 8},
			},
		},
		{name: "Gig: The Tyger and The Takimura",
			duration: 10,
			npcs: []NPC{
				{name: "Hanako Arasaka", level: 20},
				{name: "Kaoru Fujioca", level: 15},
			},
		},
	}

	totalLevel := 0
	totalNPCs := 0
	for _, quest := range quests {
		for _, npc := range quest.npcs {
			totalLevel += npc.level
			totalNPCs++
		}
	}
	averageLevel := float64(totalLevel) / float64(totalNPCs)
	fmt.Printf("Average NPC level : %.2f", averageLevel)
}
