package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "mysupersecretpassword"
	hashed := sha256.Sum256([]byte(password))
	fmt.Printf("Password: %s\n", password)
	// formato hexadecimal %x
	fmt.Printf("Hashed: %x\n", hashed)

}
