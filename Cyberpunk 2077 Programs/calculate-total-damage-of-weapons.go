package main

import "fmt"

type Weapon struct {
	name   string
	damage int
}

func main() {
	// sample weapon data
	weapons := []Weapon {
		{name: "Arasaka JKE-X2 Kenshin", damage: 86},
		{name: "Byria M-10F Peloton", damage: 112},
		{name: "Militech Crusher", damage: 135},
	}
	totalDamage := 0
	for _, weapon := range weapons {
		totalDamage += weapon.damage 
	}

	fmt.Printf("Total damage done: %d\n", totalDamage)

}
