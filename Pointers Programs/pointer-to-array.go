package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [3]int{1, 2, 3}
	ptr := &arr[0]

	for i := 0; i < len(arr); i++ {
		*(*int)(unsafe.Pointer(ptr))++
		ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])))
	}

	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Array: [%d, %d, %d]\n", arr[0], arr[1], arr[2])
}

