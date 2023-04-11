package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current working directory.
	dirName, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	for _, file := range fileInfo {
		if !file.IsDir() {
			count++
		}
	}

	fmt.Printf("There are %d files in the directory %s\n", count, dirName)
}



