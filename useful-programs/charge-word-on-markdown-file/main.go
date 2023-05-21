package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func replaceWordInMarkdown(inputFile string, outputFile string, oldWord string, newWord string) {
	// leemos el archivo de entrada
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Replace word
	updatedContent := strings.ReplaceAll(string(content), oldWord, newWord)

	// Write updated content to output markdown file
	err = ioutil.WriteFile(outputFile, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("El Archivo Markdwn Se Creo Exitosamente!")
}

func main() {
	inputFile := "input.md"
	outputFile := "output.md"
	oldWord := "burnout"
	newWord := "Neuromancer"

	replaceWordInMarkdown(inputFile, outputFile, oldWord, newWord)
}
