package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sortable is an interface for types that can be sorted.
type Sortable interface {
	Len() int         // Len returns the number of elements in the collection.
	Less(i, j int) bool // Less reports whether the element with index i should sort before the element with index j.
	Swap(i, j int)    // Swap swaps the elements with indexes i and j.
}

// IntSlice is a type representing a slice of integers.
type IntSlice []int

// Len returns the number of elements in the IntSlice.
func (s IntSlice) Len() int {
	return len(s)
}

// Less reports whether the element at index i should sort before the element at index j.
func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap swaps the elements at indexes i and j in the IntSlice.
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// quicksort is a recursive function that performs the Quick Sort algorithm on a Sortable collection.
func quicksort(s Sortable, low, high int) {
	if low < high {
		// Partition the collection and get the pivot index.
		p := partition(s, low, high)
		
		// Recursively sort the elements on the left and right sides of the pivot.
		quicksort(s, low, p-1)
		quicksort(s, p+1, high)
	}
}

// partition rearranges the elements in the collection so that elements less than the pivot
// are on the left and elements greater than the pivot are on the right.
func partition(s Sortable, low, high int) int {
	pivot := high
	i := low - 1
	
	// Iterate through the elements from low to high.
	for j := low; j < high; j++ {
		if s.Less(j, pivot) {
			// If the element at j is less than the pivot, swap it with the element at i+1.
			i++
			s.Swap(i, j)
		}
	}
	
	// Swap the pivot element with the element at i+1 to put the pivot in its correct position.
	s.Swap(i+1, pivot)
	
	// Return the index of the pivot element.
	return i + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ints := make(IntSlice, 10)
	
	// Generate random integers and store them in the IntSlice.
	for i := range ints {
		ints[i] = rand.Intn(100)
	}
	
	fmt.Println("Unsorted:", ints)
	
	// Sort the IntSlice using the quicksort algorithm.
	quicksort(ints, 0, ints.Len()-1)
	
	fmt.Println("Sorted:", ints)
}

