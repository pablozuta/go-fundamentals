package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// READ A FILE
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	
	// WRITE TO A FILE
	err = ioutil.WriteFile("test2.txt", []byte("Hello Golang"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}