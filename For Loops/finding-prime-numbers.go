package main

import "fmt"

func main() {
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d is Prime\n", i)
		}
	}
	fmt.Println()
}
