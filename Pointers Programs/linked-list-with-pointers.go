// This program implements a singly linked list in Go using pointers. It defines a Node struct with a value field and a next field that is a pointer to another Node. It also defines a LinkedList struct with a head field that is a pointer to the first Node in the list. It provides methods to add a new node to the front or end of the list, remove a node from the front or end of the list, and print the entire list.
package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) addFront(value int) {
	node := Node{value: value}
	node.next = list.head
	list.head = &node
}

func (list *LinkedList) addEnd(value int) {
	node := Node{value: value}

	if list.head == nil {
		list.head = &node
		return
	}
	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = &node
}

func (list *LinkedList) removeFront() {
	if list.head == nil {
		return
	}
	list.head = list.head.next
}

func (list *LinkedList) removeEnd() {
	if list.head == nil {
		return
	}
	current := list.head

	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (list *LinkedList) print() {
	current := list.head

	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	list.addFront(1)
	list.addFront(2)
	list.addFront(3)
	list.addEnd(4)
	list.addEnd(5)
	list.addEnd(6)

	list.print()

	list.removeFront()
	list.removeEnd()

	list.print()
}
