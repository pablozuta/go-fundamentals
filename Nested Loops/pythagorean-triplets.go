// A Pythagorean triple consists of three positive integers a, b, and c, such that a2 + b2 = c2. Such a triple is commonly written (a, b, c), and a well-known example is (3, 4, 5).
package main

import "fmt"

func main() {
	limit := 20
	for a := 1; a <= limit; a++ {
		for b := a; b <= limit; b++ {
			for c := b; c <= limit; c++ {
				if (a*a)+(b*b) == c*c {
					fmt.Printf("%d %d %d\n", a, b, c)
				}
			}
		}
	}
}
