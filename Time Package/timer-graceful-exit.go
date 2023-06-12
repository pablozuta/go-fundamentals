package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	duracion := 10 * time.Second
	timer := time.NewTimer(duracion)

	fmt.Println("El tiempo empezo a correr")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <- timer.C:
		fmt.Println("El tiempo termino")
		
	case <- signals:
		fmt.Println("Tiempo Interrumpido")
		if !timer.Stop() {
			<- timer.C
		}

	}
}