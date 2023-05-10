package main

import "fmt"

func main() {
	source := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	dest := make([]byte, 3)
	copy(dest, source[2:5])
	fmt.Printf("Source %v\n", source)
	fmt.Printf("%v", dest)
}