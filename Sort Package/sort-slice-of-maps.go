package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []map[string] int {
		{"Alice": 25},  
        {"Bob": 30},  
        {"Charlie": 20},  
        {"David": 30},  
        {"Eve": 25},
	}
	// sort the map by value
	sort.Slice(ages, func(i, j int) bool {
		for _, age := range ages[i] {
			if age != ages[j][getMapKey(ages[i], age)] {
				return age < ages[j][getMapKey(ages[i], age)]
			}
		}
		return false
	})
	// print the sorted map
	fmt.Println(ages)
	
}

func getMapKey(m map[string]int, val int)string  {
	for key, value := range m {
		if value == val {
			return key
		}
	}
	return ""
}