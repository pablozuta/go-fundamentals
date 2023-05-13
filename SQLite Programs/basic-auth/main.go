package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// open the database
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ask the user for username and password
	var username string
	var password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	// query the database for the user
	var storedPassword string
	err = db.QueryRow("SELECT password FROM userstable WHERE username =?", username).Scan(&storedPassword)
	if err != nil {
		log.Println("*****Password or user Incorrect, Goodbye!*****")
		return
	}

	// compare the passwords
	if password == storedPassword {
		// show the menu
		fmt.Println("Welcome, ", username)
		for {
			fmt.Println("Please choose an option:")
			fmt.Println("1. Print 'Hello Golang'")
			fmt.Println("2. Print the sum of two numbers")
			fmt.Println("3. Print the factorial of a number")
			fmt.Println("4. Exit")
			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				fmt.Println("Hello Golang")
			case 2:
				var a, b int
				fmt.Print("Enter the first number :")
				fmt.Scanln(&a)
				fmt.Print("Enter the second number :")
				fmt.Scanln(&b)
				fmt.Println("The sum is ", a+b)
			case 3:
				var n int
				fmt.Print("Enter the number: ")
				fmt.Scanln(&n)
				f := 1
				for i := 2; i < n; i++ {
					f *= i
				}
				fmt.Println("The factorial is: ", f)
			case 4:
				fmt.Println("Goodbye! :)")
				return
			default:
				fmt.Println("Invalid choice")

			}

		}
	} else {
		fmt.Println("Invalid username or password")
	}
}
