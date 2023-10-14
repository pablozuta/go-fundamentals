// This program is an "Ethical Password Strength Checker" that assesses the strength of a user's password. It encourages ethical practices by emphasizing the importance of strong, secure passwords.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the ethical password strength checker")
	fmt.Print("Enter your password: ")

	// creamos una variable donde guardaremos el valor recibido
	var password string
	fmt.Scan(&password)

	// guardamos el resultado de la funcion en una variable
	strength := checkPasswordStrenght(password)
	// mostramos el resultado por pantalla
	fmt.Printf("Password Strength: %s\n", strength)
}

func checkPasswordStrenght(password string) string {
	if len(password) < 8 {
		return "Weak (too short)"
	}
	// Additional ethical checks can be added here, e.g., checking for common passwords, enforcing special characters, etc.

	return "Strong"
}
