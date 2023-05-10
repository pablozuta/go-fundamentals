package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open the csv file
	file, err := os.Open("pokemon_data.csv")
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

	// print the number of records in the file
	fmt.Printf("Number of records = %d\n", len(records))

	// calculate the average value of a column
	sum := 0.0
	for _, record := range records {
		value := record[6]
		if value != "" {
			sum += parseFloat(value)

		}
	}
	avg := sum / float64(len(records))
	fmt.Printf("Average value: %.2f\n", avg)

}

func parseFloat(str string) float64 {
	// Parse a string into a float64 value
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return value
}
