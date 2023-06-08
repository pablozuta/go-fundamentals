package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int {3, 7, 4, 1, 6, 5, 3, 2, 1, 8, 3,}
	sort.Ints(nums)
	fmt.Println(nums)
}