package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	message := []byte("Hello Golang")
	hashed := sha256.Sum256(message)

	signature, err :=  rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message %s\n", message)
	fmt.Printf("Signature %x\n", signature)
}