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
	// print every record
	for _, record := range records {
		fmt.Println(record)
	}

	// ask user for input
	var country string
	fmt.Print("Enter a country:")
	fmt.Scan(&country)


	// print each record filtering by user input
	count := 0
	for _, record := range records {
		if record[3] == country {
			fmt.Println(record)
            count ++
		}
	}
	fmt.Printf("%v registros para %s\n", count, country)

}