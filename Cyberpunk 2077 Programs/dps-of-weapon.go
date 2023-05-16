package main

import "fmt"

type WeaponTwo struct {
	name       string
	damage     int
	rateOfFire int
}

func main() {
	// sample weapon data
	weapons := []WeaponTwo{
		{name: "Arasaka JKE-X2 Kenshin", damage: 86, rateOfFire: 2},
		{name: "Burya M-10F Peloton", damage: 112, rateOfFire: 3},
		{name: "Militech Crusher", damage: 135, rateOfFire: 1},
	}
	totalDamage := 0
	totalRateOfFire := 0
	for _, weapon := range weapons {
		totalDamage += weapon.damage
		totalRateOfFire += weapon.rateOfFire
	}
	averageDPS := float64(totalDamage) / float64(totalRateOfFire)
	fmt.Printf("Average DPS: %.2f\n", averageDPS)

}
