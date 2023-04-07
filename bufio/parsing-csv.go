package main

import (
	"fmt"
	"bufio"
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Open("pokemon_data.csv")
	if err != nil {
		fmt.Println("no se pudo leer archivo")
		return
	}
	
	defer file.Close()
	
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println("no se pudo leer archivo")
			break

	}
	fmt.Println(record)
	}

}