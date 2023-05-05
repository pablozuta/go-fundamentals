// This program takes an array of cyberpunk anime titles as input and a search string, and then searches for the search string in the array using a binary search algorithm.
package main

import "fmt"

func binarySearch(arr []string, search string) int {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == search {
			return i
		}
	}
	return -1
}

func main() {
	titles := []string{
		"Akira",
		"Ghost in the Shell",
		"Psycho-Pass",
		"Serial Experiments Lain",
		"Ergo Proxy",
		"Texhnolyze",
		"Blade Runner: Black Lotus",
		"Battle Angel Alita",
		"Armitage III",
		"Cyber City Oedo 808",
	}

	search := "Ghost in the Shell"
	index := binarySearch(titles, search)

	if index == -1 {
		fmt.Println("Title not found")
	} else {
		fmt.Printf("'%s' found at index %d in the array.", search, index)
	}

}
