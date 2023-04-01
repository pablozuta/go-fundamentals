package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// READ A DIRECTORY
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	
	// CREATING A TEMPORARY FILE AND DIRECTORY
	tempFile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()
	fmt.Println("Created temporary file", tempFile.Name())
	
	tempDir, err :=ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)
	fmt.Println("Created temporary directory", tempDir)
	
}
