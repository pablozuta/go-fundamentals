package main

import (
	"fmt"
	"sort"
)

type Person3 struct {
	Name string
	Age  int
}
type ByAge []Person3

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person3{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
		{"David", 30},
		{"Eve", 25},
	}
	// sort the people by age
	sort.Sort(ByAge(people))

	// show the results
	fmt.Println(people)

}
