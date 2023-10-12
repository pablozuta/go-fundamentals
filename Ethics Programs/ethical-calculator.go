// This program is an "Ethical Calculator" that prompts the user for two numbers and an operation and performs the calculation while emphasizing ethical considerations, such as avoiding division by zero.
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Ethical Calculator!!")
	fmt.Println("This program allows you to perform arithmetic operations between two numbers")
	// creamos variable para guardar los valores de dos numeros y un operador aritmetico
	var num1, num2 float64
	var op string

	// input para el primer numero
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Invalid input for the first number", err)
		return
	}
	// input para el segundo numero
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid input for the second number.", err)
		return
	}
	// input para el operador aritmetico
	fmt.Print("Enter the arithmetic operator (+ - * /): ")
	_, err = fmt.Scan(&op)
	if err != nil || (op != "+" && op != "-" && op != "*" && op != "/") {
		fmt.Println("Invalid input for operator:", err)
		return
	}

	result, _ := calculate(num1, num2, op)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("The result of the arithmetic operation is:", result)
}

func calculate(x, y float64, op string) (float64, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, errors.New("You cannot divide by zero")
		}
		return x / y, nil
	default:
		return 0, errors.New("Not a valid operation")

	}
}
