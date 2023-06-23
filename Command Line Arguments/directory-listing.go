package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	// read the content of the directory
	files, err := ioutil.ReadDir(args[1])
	if err != nil {
		panic(err)
	}
	// mostrar los nombres de los directorios
	for _, file := range files {
		fmt.Println(file.Name())
	}
}