package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tealeg/xlsx"
)

func main() {
	// abre una coneccion con la base de datos
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
		fmt.Println("4. EXPORTAR EN FORMATO EXCEL")
		fmt.Println("5. SALIR")
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
				fmt.Fprintf(os.Stderr, "Error al borrar usuario: %v\n", err)
			} else {
				rowsAffected, _ := result.RowsAffected()
				if rowsAffected == 0 {
					fmt.Println("Usuario no encontrado.")
				} else {
					fmt.Println("User Borrado Exitosamente.")
				}
			}
		case 3:
			rows, err := db.Query("SELECT id, name, email FROM users")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error al mostrar usuarios: %v\n", err)
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
			// Consulta la base de datos para obtener la tabla de usuarios
			rows, err := db.Query("SELECT * FROM users")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error al consultar usuarios: %v\n", err)
			} else {
				// crea un nuevo archivo excel
				file := xlsx.NewFile()
				sheet, err := file.AddSheet("Usuarios")
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error al agregar hoja: %v\n", err)
				} else {
					// agrega encabezados de columna
					row := sheet.AddRow()
					row.AddCell().SetValue("ID")
					row.AddCell().SetValue("Nombre")
					row.AddCell().SetValue("Email")

					// itera a través de las filas de resultados y agrega datos a las celdas
					for rows.Next() {
						var id int
						var name string
						var email string
						rows.Scan(&id, &name, &email)

						row := sheet.AddRow()
						row.AddCell().SetValue(id)
						row.AddCell().SetValue(name)
						row.AddCell().SetValue(email)
					}

					// guarda el archivo y muestra mensaje de éxito
					err = file.Save("usuarios.xlsx")
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error al guardar archivo: %v\n", err)
					} else {
						fmt.Println("Tabla de Usuarios Exportada a 'usuarios.xlsx'.")
					}
				}
			}
		case 5:
			fmt.Println("HASTA LUEGO!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")

		}

		fmt.Println()
	}
}
