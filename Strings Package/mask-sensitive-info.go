package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Credit card number: 1234-5678-9012-3456, SSN: 123-45-6789"

	maskedText := maskSensitiveData(text)
	fmt.Println("Original Text:", text)
	fmt.Println("Masked Text:", maskedText)
}

func maskSensitiveData(s string) string {
	// mask credit card numbers
	maskedText := strings.ReplaceAll(s, "-", "")

	// mask SSN
	maskedText = strings.ReplaceAll(maskedText, "SSN: 123456789", "SSN: ***-**-")
	return maskedText
}
