package main

import (
	"fmt"
	"math/rand"
	"time"
)

var horoscopes = []string{
	"Today is a great day to take risks and try something new!",
	"You will encounter unexpected opportunities today. Embrace them!",
	"Take some time for self-reflection today. It will help you grow.",
	"Your hard work and dedication will pay off today. Stay focused!",
	"Avoid making hasty decisions today. Take your time and think things through.",
}

func generateHoroscope() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(horoscopes))
	return horoscopes[index]
}

func main() {
	horoscope := generateHoroscope()
	fmt.Println("Your horoscope for today:")
	fmt.Println(horoscope)
}
