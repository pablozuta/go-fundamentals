// In this program, we define a Stack data structure using the [T any] syntax to indicate that it is generic and can hold elements of any type. The Push method takes an element of type T and appends it to the end of the elements slice. The Pop method removes and returns the last element in the elements slice.

package main

import "fmt"

// definimos un struct que puede inicializarse con any data type
type Stack[T any] struct {
	// declaramos un slice de tipo T
	// The type parameter T is a placeholder for a specific type that will be determined when the Stack is instantiated with an actual type.
	elements []T
}

// In the Push method definition, the receiver of the method is a pointer to a Stack[T] struct, which is denoted by s. The Stack[T] struct is the object that the method is called on, and it provides access to the elements field, which is a slice that stores the elements in the stack.
// Inside the Push method, the elements field is the receiver of the append function call, which is a built-in function in Go that can be used to append elements to a slice. The append function takes two arguments: the slice to append to, and the element to append.
// In this case, the slice to append to is s.elements, which is the slice of elements stored in the Stack[T] struct. The second argument is the element parameter passed to the Push method, which is of the same type as the type parameter T.
// When the append function is called, it appends the element argument to the end of the s.elements slice, and returns a new slice that contains the old elements plus the new element. The Push method then assigns the new slice to the s.elements field, which updates the contents of the slice stored in the Stack[T] struct.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	// if the elements field of the Stack[T] struct is empty the method panics with an error message.
	if len(s.elements) == 0 {
		panic("Stack is empty!")
	}
	// This line retrieves the last element of the elements field of the Stack[T] struct using the index len(s.elements)-1. This is the top element of the stack.
	element := s.elements[len(s.elements)-1]
	// This line removes the top element from the elements field of the Stack[T] struct by slicing the slice to exclude the last element.
	s.elements = s.elements[:len(s.elements)-1]
	// This line returns the top element that was retrieved earlier.
	return element
}

func main() {
	// In the main function, we create two Stack instances with different types and push some elements onto them. We then call the Pop method to retrieve the elements in reverse order.

	// creamos un stack de integers y agregamos elementos
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("intStack", intStack)
	fmt.Println(intStack.Pop()) // Output: 3
	fmt.Println(intStack.Pop()) // Output: 2
	fmt.Println(intStack.Pop()) // Output: 1
	fmt.Println("intStack", intStack)
	fmt.Println()

	// creamos un stack de strings y agregamos elementos
	stringStack := Stack[string]{}
	stringStack.Push("Pedro Paramo")
	stringStack.Push("El LLano en LLamas")

	fmt.Println("stringStack", stringStack)
	fmt.Println(stringStack.Pop()) // Output: "El LLano en LLamas"
	fmt.Println(stringStack.Pop()) // Output: "Pedro Paramo"
	fmt.Println("stringStack", stringStack)

}
