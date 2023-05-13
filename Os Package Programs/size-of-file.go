package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "archivo.txt"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := fileInfo.Size()
	fmt.Printf("el tamaño del archivo %s is %d bytes.", fileName, fileSize)
}