// In this example, we define a custom error type NegativeNumberError that includes the negative number encountered. The Error method is implemented for this type to provide a custom error message.

// The divide function now checks if either a or b is negative, and if so, returns a NegativeNumberError with the corresponding negative number. In the main function, we handle this custom error type separately and print the custom error message.
package main

import (
	"errors"
	"fmt"
)

// Define a custom error type for negative numbers.
type NegativeNumberError struct {
	number int
}

// Implement the Error() method for the custom error type.
func (e *NegativeNumberError) Error() string {
	return fmt.Sprintf("Negative Number NOT Allowed %d", e.number)
}

// Function to perform division and handle errors.
func divide(a, b int) (int, error) {
	// Check for division by zero.
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	// Check for negative numbers in the input.
	if a < 0 || b < 0 {
		// Return a custom error with the negative number.
		return 0, &NegativeNumberError{number: a}
	}
	// Perform the division and return the result.
	return a / b, nil
}

func main() {
	// Example 1: Performing a valid division.
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Attempting division by zero.
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 3: Handling the custom error for negative numbers.
	result, err = divide(-10, 2)
	if err != nil {
		if e, ok := err.(*NegativeNumberError); ok {
			fmt.Println("Negative Number Error:", e)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Result:", result)
	}
}
