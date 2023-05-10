package main

import "fmt"

func main() {
	bytes := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	fmt.Printf("Normal %v\n", bytes)
	reverseBytes(bytes)
	fmt.Printf("Reversed %v", bytes)
}

func reverseBytes(bytes []byte)  {
	for i := 0; i < len(bytes) / 2; i++ {
		j := len(bytes) - i - 1
		bytes[i], bytes[j] = bytes[j], bytes[i]
		
	}
	
}