package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	data := "Hola Golang"
	hash := sha256.Sum256([]byte(data))
	hashString := hex.EncodeToString(hash[:])
	fmt.Printf("%s => %s", data, hashString)
}