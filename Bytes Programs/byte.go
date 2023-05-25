package main

import "fmt"

func main() {
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
    dataDos := []byte {0x32, 0x41, 0x37}


	fmt.Println(string(data)) // output hello
	fmt.Println(string(dataDos)) // output 2A7

}