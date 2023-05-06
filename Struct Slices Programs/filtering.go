package main

import "fmt"

type JazzRecord struct {
	Title string
	Year  int
}

func main() {
	records := []JazzRecord{
		{Title: "Giant Steps", Year: 1960},
		{Title: "Kind of Blue", Year: 1959},
		{Title: "Mingus Ah Um", Year: 1959},
	}

	fmt.Println("Original records", records)

	var matchingRecords []JazzRecord

	for _, record := range records {
		if record.Year == 1960 {
			matchingRecords = append(matchingRecords, JazzRecord{Title: record.Title, Year: record.Year})
		}
	}
	fmt.Println("Matching Records: ", matchingRecords)
}
