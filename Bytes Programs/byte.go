package main

import "fmt"

func main() {
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}

	fmt.Println(string(data)) // output hello

}