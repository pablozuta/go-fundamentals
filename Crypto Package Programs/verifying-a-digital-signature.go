package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Message to be signed
	message := []byte("Hello Digital")

	// Hash the message using SHA256
	hashed := sha256.Sum256(message)

	// Sign the hashed message with the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	// Extract the public key from the private key
	publicKey := privateKey.PublicKey

	// Verify the signature using the public key
	err = rsa.VerifyPKCS1v15(&publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signature Verificada Exitosamente.")
}
