package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("new-file.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	defer file.Close()
	
	_, err = file.WriteString("Black Music")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("data escrita exitosamente.")
	}
}