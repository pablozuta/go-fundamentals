// This program takes an array of strings as input and a search string, and then searches for the search string in the array using a linear search algorithm.
package main

import "fmt"

func linearSearch(arr[]string, search string) int  {
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			return i
		}
	}
	return -1
}

func main()  {
	arr := [] string {"manzana", "pera", "sandia", "uva", "durazno"}
	search := "uva"
	index := linearSearch(arr, search)

	if index == -1 {
		fmt.Printf("String %s no encontrada en array.\n", search)
	} else {
		fmt.Printf("String %s encontrada en index[%d]",search, index)
	}
	
}