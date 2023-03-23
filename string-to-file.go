package main

import (
	"io/ioutil"
	"log"
)

func main() {
	myString := "La musica de tus pasos"
	
	err:= ioutil.WriteFile("output.txt", []byte(myString), 0644)

	if err != nil {
		log.Fatal(err)
	}
}