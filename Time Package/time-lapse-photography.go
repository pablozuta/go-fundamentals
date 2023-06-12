package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	duration := 2 * time.Minute
	tick := time.Tick(1 * time.Minute)

	fmt.Println("Comienzo de fotografia time-lapse")

	// create output directory
	outputDir := "output"
	err := os.Mkdir(outputDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error al crear directorio", err)
		return
	}
	// Capture photos
	for i := 0; ; i++ {
		select {
		case <-time.After(duration):
			fmt.Println("Time's up!")
			return
		case <-tick:
			// Capture photo
			filename := filepath.Join(outputDir, time.Now().Format("2006-01-02-15-04-05.jpg"))
			fmt.Println("Capturing photo:", filename)

			// Simulate photo capture time
			time.Sleep(5 * time.Second)

			// Save photo
			f, err := os.Create(filename)
			if err != nil {
				fmt.Println("Error saving photo:", err)
				return
			}
			f.Close()
		}
	}
}

func handleInterrupt() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	fmt.Println("Interrupted")
	os.Exit(0)
}
