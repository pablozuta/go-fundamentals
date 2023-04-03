package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 5 * time.Second
	timer := time.NewTimer(duration)

	fmt.Println("El tiempo empezo a correr")
	
	<-timer.C

	fmt.Println("Pasaron", duration)
}