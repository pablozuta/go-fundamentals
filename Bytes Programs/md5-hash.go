package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "Hello World"
	hash := md5.Sum([]byte(str))
	fmt.Printf("%x", hash)
}