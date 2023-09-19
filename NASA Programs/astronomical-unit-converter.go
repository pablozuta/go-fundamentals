// This program converts distances from Astronomical Units (AU) to kilometers and miles. It defines constants for the conversion factors and provides functions to convert AU to kilometers and miles. You can input a distance in AU, and the program will calculate and display the equivalent distance in kilometers and miles.

package main

import "fmt"

const (
	OneAUInKilometers = 149_597_870.7 // One Astronomical Unit in kilometers
	OneAUInMiles      = 92_955_807.27 // One Astronomical Unit in miles
)

func main() {
	distanceInAU := 5.2 // Example distance in Astronomical Units

	distanceInKilometers := AUToKilometers(distanceInAU)
	distanceInMiles := AUToMiles(distanceInAU)

	fmt.Printf("%.2f Astronomical Units are Equal to:\n", distanceInAU)
	fmt.Printf("%.2f Kilometers\n", distanceInKilometers)
	fmt.Printf("%.2f Miles\n", distanceInMiles)
}

func AUToKilometers(au float64) float64 {
	return au * OneAUInKilometers
}

func AUToMiles(au float64) float64 {
	return au * OneAUInMiles
}
