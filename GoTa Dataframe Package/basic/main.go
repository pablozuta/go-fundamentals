package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	// Create a series (column) for the data frame
	col1 := series.New([]string{"John", "Jane", "Bob"}, series.String, "Name")
	col2 := series.New([]int{25, 30, 35}, series.Int, "Age")

	// Create a data frame
	df := dataframe.New(col1, col2)

	// Print the data frame
	fmt.Println(df)

	// Accessing values in the data frame
	fmt.Println(df.Col("Name").Records())
	fmt.Println(df.Col("Age").Records())
}

