package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// leyendo desde standard input
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("no se pudo leer desde standard input")
		return
	}
	if err != nil {
		fmt.Println("no se pudo leer desde standard input")
		return
	}
	fmt.Println("leyendo desde standard input: ", string(data))
	
	// escribir a standard output
	err = ioutil.WriteFile(os.Stdout.Name(), []byte("Fair Trade\n"), 0644)
	if err != nil {
		fmt.Println("no se pudo escribir a standard output: ", err)
		return
	}

}