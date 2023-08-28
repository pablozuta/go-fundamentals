package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter commands for the Mars Rover: ")
	if scanner.Scan() {
		commands := scanner.Text()
		executeCommands(commands)
	}
}

func executeCommands(commands string)  {
	fmt.Printf("Executing commands: %s\n", commands)
	fmt.Println("Mission Finished Succesfully!!")
}