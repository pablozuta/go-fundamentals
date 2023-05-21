package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error al abrir archivo.", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error al crear archivo.", err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into phrases using a dot (".") delimiter
		phrases := strings.Split(line, ".")

		// Shuffle the phrases randomly
		shuffledPhrases := shuffle(phrases)

		// Join the shuffled phrases back into a line using a dot (".") delimiter
		outputLine := strings.Join(shuffledPhrases, ".") + "\n"

		// Write the shuffled line to the output file
		_, err := outputFile.WriteString(outputLine)
		if err != nil {
			fmt.Println("Error al escribir en archivo de salida.", err)
			return
		}
	}

	fmt.Println("Archivo escrito exitosamente.")
}

// Shuffle the phrases randomly using the Fisher-Yates algorithm
func shuffle(phrases []string) []string {
	shuffled := make([]string, len(phrases))
	perm := rand.Perm(len(phrases))
	for i, v := range perm {
		shuffled[v] = phrases[i]
	}
	return shuffled
}
