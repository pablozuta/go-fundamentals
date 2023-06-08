package main

import (
	"fmt"
	"unsafe"
)

type Person2 struct {
	Name string
	Age int
}

func main() {
	p := Person2 {"John", 32}

	pAddr := unsafe.Pointer(&p)

	fmt.Printf("Memory use of person struct:\n")
	fmt.Printf("\tAddress: %p\n", &pAddr)
	fmt.Printf("\tSize: %d bytes\n", unsafe.Sizeof(p))
	fmt.Printf("\tName field offset: %d bytes\n", unsafe.Offsetof(p.Name))
	fmt.Printf("\tAge field offset: %d bytes\n", unsafe.Offsetof(p.Age))
}