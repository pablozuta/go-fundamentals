package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// define the message to sign
	message := []byte("Cyberpunk 2077")
	// compute the SHA-256 version of the message
	hash := sha256.Sum256(message)
	// sign the hash using the RSA private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}
	// verify the signature using the RSA public key
	publicKey := &privateKey.PublicKey
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		fmt.Println("ERROR al verificar signature", err)
	} else {
		fmt.Println("ACCION EXITOSA")
	}

}
