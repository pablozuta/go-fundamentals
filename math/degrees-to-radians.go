package main

import (
	"fmt"
	"math"
)

func main() {
	degrees := 45.0
	radians := degrees * math.Pi / 180
	fmt.Printf("%.2f degrees is equivalent to %.2f radians", degrees, radians)

}