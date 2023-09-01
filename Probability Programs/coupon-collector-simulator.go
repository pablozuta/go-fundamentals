package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const totalCoupons = 50
	const numSimulations = 10000

	var collectedCoupons [totalCoupons]bool

	for i := 0; i < numSimulations; i++ {
		coupons := make(map[int]bool)
		collectedCount := 0

		for {
			coupon := rand.Intn(totalCoupons)
			if !coupons[coupon] {
				coupons[coupon] = true
				collectedCount++

				if collectedCount == totalCoupons {
					break
				}
			}
		}
		collectedCoupons[collectedCount-1] = true
	}
	probability := float64(numSimulations) / float64(len(collectedCoupons))
	fmt.Printf("Probability of collecting all %d coupons: %.4f\n", totalCoupons, probability)

}
