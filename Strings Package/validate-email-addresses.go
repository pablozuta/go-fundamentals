package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{
		"john@example.com",
		"jane.doe@example",
		"invalid-email",
		"contact@example.com",
	}
	fmt.Println("Validating Email Addresses:")
	for _, email := range emails {
		if isValidEmail(email) {
			fmt.Printf("%s is a valid email address\n", email)
		} else {
			fmt.Printf("%s is NOT a valid email address\n", email)
		}
	}
}

func isValidEmail(email string) bool {
	// check if the email contains a single @
	if strings.Count(email, "@") != 1 {
		return false
	}
	// split the email into local part and domain part
	parts := strings.Split(email, "@")
	localPart, domainPart := parts[0], parts[1]

	// check if the local part is not empty
	if len(localPart) == 0 {
		return false
	}
	// check if the domain part is not empty
	if len(domainPart) == 0 {
		return false
	}
	// check if the domain part has at least one dot
	if strings.Index(domainPart, ".") == -1 {
		return false
	}
	return true

}
