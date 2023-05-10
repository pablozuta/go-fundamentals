package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func main() {
	// valores guardados en un slice de bytes (In Go, a byte is an alias for the uint8 type)
    key := []byte("mysupersecretkey") // 16 bits
	fmt.Printf("Valor en bytes de key %v, longitud %v\n",key, len(key) )
    message := []byte("hello, golang crypto!")
	// In cryptography, a nonce (which stands for "number used once") is a random or pseudo-random value that is used only once in a cryptographic communication to ensure the freshness or uniqueness of the communication.
    nonce := make([]byte, 12) // 12 bits inicializados a su zero-value

    // Create a new AES cipher block using the key
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
	

    // Create a new GCM cipher
	// The NewCipher function takes a byte slice key as its argument and returns a new cipher.Block interface that can be used to perform AES encryption and decryption operations. The key byte slice must be 16, 24, or 32 bytes long to create a valid AES cipher block for 128-bit, 192-bit, or 256-bit key sizes, respectively.
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err)
    }

    // Generate a random nonce
	// This code is using the io.ReadFull() function to read random bytes from the rand.Reader source and fill the nonce byte slice with those random bytes. The rand.Reader is a cryptographic random number generator provided by the Go standard library, which is designed to produce high-quality random bytes suitable for use in cryptographic applications.
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err)
    }
	fmt.Printf("Value of nonce after generation: %v\n", nonce) // cada vez que corramos el programa este valor sera diferente

    // Encrypt the message using the GCM cipher(esto genera un byte slice)
    ciphertext := aesgcm.Seal(nil, nonce, message, nil)
    fmt.Printf("Ciphertext en formato byte slice: %v, longitud=%v\n", ciphertext, len(ciphertext))
	// en formato hexadecimal
    fmt.Printf("Ciphertext: %x\n", ciphertext)

    // Decrypt the ciphertext using the GCM cipher
    plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Plaintext: %s\n", plaintext)
}
