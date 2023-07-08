package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	content, err := readFile("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Content:", content)
	}

}
