// This program uses the Int and Bool functions to define two command line flags: length and symbols. It then uses flag.Parse() to parse the command line arguments.
// The program generates a random password of the specified length, using only letters and numbers by default. If the symbols flag is set to true, it includes additional symbols in the password.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// define command-line flags
	lengthPtr := flag.Int("length", 12, "length of password")
	symbolsPtr := flag.Bool("symbols", false, "include symbols")

	// parse command-line flags
	flag.Parse()

	// generate random seed
	rand.Seed(time.Now().UnixNano())

	// generate password
	var characters string
	if *symbolsPtr {
		characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	} else {
		characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	password := make([]byte, *lengthPtr)
	for i := range password {
		password[i] = characters[rand.Intn(len(characters))]
	}

	// print password
	fmt.Println(string(password))
}
