package main

import "fmt"

func moveDisc(numDiscs int, fromPole string, toPole string, viaPole string) {
	// base case: if there is only one disc, move it from the fromPole to the toPole
	if numDiscs == 1 {
		fmt.Printf("Move disc 1 from %s to %s\n", fromPole, toPole)
		return
	}
	// move n-1 discs from the fromPole to the viaPole
	moveDisc(numDiscs-1, fromPole, viaPole, toPole)

	// move the nth disc from the fromPole to the toPole
	fmt.Printf("Move disc %d from %s to %s\n", numDiscs, fromPole, toPole)
	// move n-1 discs from the viaPole to the toPole
	moveDisc(numDiscs-1, viaPole, toPole, fromPole)

}
func main() {
	numDiscs := 3
	moveDisc(numDiscs, "A", "C", "B")
}
