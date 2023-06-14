package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// iniciamos semilla random
	rand.Seed(time.Now().UnixNano())

	deck := generateDeck()
	hand := drawHand(deck, 5)

	favorableOutcomes := countFlushPossibilities(deck, hand)
	probability := float64(favorableOutcomes) / float64(len(deck)-5)

	// mostramos la mano de cartas
	for _, card := range hand {
		fmt.Println(card)
	}
	// mostramos la probabilidad de un flush
	fmt.Printf("Probability of getting a flush in a 5-card poker hand: %.4f", probability)

}

func generateDeck() []string {
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	deck := make([]string, 52)
	index := 0

	for _, suit := range suits {
		for _, rank := range ranks {
			card := rank + " of " + suit
			deck[index] = card
			index++
		}
	}
	return deck
}

func drawHand(deck []string, numCards int) []string {
	hand := make([]string, numCards)

	for i := 0; i < numCards; i++ {
		randomIndex := rand.Intn(len(deck))
		hand[i] = deck[randomIndex]
		deck = append(deck[:randomIndex], deck[randomIndex+1:]...)
	}
	return hand
}

func checkFlush(hand []string) bool {
	suits := make(map[string]bool)

	for _, card := range hand {
		suit := card[strings.Index(card, "of")+3:]
		suits[suit] = true
	}
	return len(suits) == 1
}
func countFlushPossibilities(deck []string, hand []string) int {
	suits := make(map[string]int)

	for _, card := range deck {
		suit := card[strings.Index(card, "of")+3:]
		suits[suit]++
	}

	for _, card := range hand {
		suit := card[strings.Index(card, "of")+3:]
		suits[suit]--
	}

	count := 0
	for _, suitCount := range suits {
		if suitCount >= 5 {
			count++
		}
	}

	return count
}
