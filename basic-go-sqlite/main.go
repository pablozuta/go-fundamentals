package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// abre una connecion con la base de datos
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// crea una tabla si no existe
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		panic(err)
	}

	// crea un loop que pregunta al usuario por input
	for {
		fmt.Println("1. AGREGAR USUARIO")
		fmt.Println("2. BORRAR USUARIO")
		fmt.Println("3. MOSTRAR USUARIOS")
		fmt.Println("4. SALIR")
		fmt.Print("INGRESA TU OPCION: ")
        
		// almacena numero en variable
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter name: ")
			var name string
			fmt.Scan(&name)

			fmt.Print("Enter email: ")
			var email string
			fmt.Scan(&email)

			_, err = db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", name, email)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error al agregar usuario: %v\n", err)
			} else {
				fmt.Println("Usuario Agregado Correctamente.")
			}
		case 2:
			fmt.Print("Enter user ID: ")
			var id int
			fmt.Scan(&id)

			result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting user: %v\n", err)
			} else {
				rowsAffected, _ := result.RowsAffected()
				if rowsAffected == 0 {
					fmt.Println("User not found.")
				} else {
					fmt.Println("User deleted successfully.")
				}
			}
		case 3:
			rows, err := db.Query("SELECT id, name, email FROM users")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error showing users: %v\n", err)
			} else {
				defer rows.Close()

				fmt.Println("ID\tNombre\tEmail")
				for rows.Next() {
					var id int
					var name string
					var email string
					err = rows.Scan(&id, &name, &email)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error showing users: %v\n", err)
						break
					}
					fmt.Printf("%d\t%s\t%s\n", id, name, email)
				}

				err = rows.Err()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error showing users: %v\n", err)
				}
			}
		case 4:
			fmt.Println("HASTA LUEGO!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}
