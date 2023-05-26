package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate a new RSA private key with a key size of 2048 bits
	// se usa el metodo rand.Reader del crypto package para generar numeros aleatorios
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// The message to be signed
	message := []byte("Hello Golang")

	// Hash the message using SHA-256
	hashed := sha256.Sum256(message)

	// Sign the hashed message using the private key and the PKCS #1 v1.5 padding scheme
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	// Print the original message and the generated signature
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Signature: %x\n", signature)
}
