package main

import "fmt"

func main() {
	map1 := map[string]int {
		"a": 1,
		"b": 2,
		"c": 3,
	}

	map2 := map[string]int {
		"d": 4,
		"e": 5,
		"f": 6,
	}
	mergedMap :=make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] = value
	}
	fmt.Println("Merged Map:", mergedMap)
}