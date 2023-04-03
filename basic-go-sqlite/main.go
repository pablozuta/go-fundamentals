package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a table if it doesn't already exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		panic(err)
	}

	// Insert a new row into the table
	_, err = db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", "Alice Coltrane", "alice@gmail.com")
	if err != nil {
		panic(err)
	}

	// Query the table for all rows
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows and print the results
	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// Check for errors during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}