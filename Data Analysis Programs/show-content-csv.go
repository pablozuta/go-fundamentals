package main

import (
	"fmt"
	"os"
	"encoding/csv"
	)

func main() {
	// open the csv file
	file, err := os.Open("arwu-iberoamerica.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// create a new csv reader
	reader := csv.NewReader(file)

	// read in all the records
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// print each record
	for _, record := range records {
		fmt.Println(record)
	}

}