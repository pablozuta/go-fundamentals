package main

import "fmt"

func main() {
	bytes1 := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	bytes2 := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	bytes3 := []byte{0x01, 0x02, 0x03, 0x04, 0x06}

	if equal(bytes1, bytes2) {
		fmt.Println("bytes1 and bytes2 are equal")
	} else {
		fmt.Println("bytes1 and bytes2 are NOT equal")
	}

	if equal(bytes1, bytes3) {
		fmt.Println("bytes1 and bytes3 are equal")
	} else {
		fmt.Println("bytes1 and bytes3 are NOT equal")
	}
}

func equal(bytes1, bytes2 []byte) bool {
	if len(bytes1) != len(bytes2) {
		return false
	}

	for i := 0; i < len(bytes1); i++ {
		if bytes1[i] != bytes2[i] {
			return false
		}
	}
	return true
}
