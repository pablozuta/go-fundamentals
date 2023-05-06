package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tealeg/xlsx"
)

func main() {
	// abre una coneccion con la base de datos
	db, err := sql.Open("sqlite3", "bookstore.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// crea una tabla si no existe
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, title TEXT, author TEXT)")
	if err != nil {
		panic(err)
	}

	// crea un loop que pregunta al usuario por input
	for {
		fmt.Println("1. AGREGAR LIBRO")
		fmt.Println("2. BORRAR LIBRO")
		fmt.Println("3. MOSTRAR LIBROS")
		fmt.Println("4. EXPORTAR EN FORMATO EXCEL")
		fmt.Println("5. SALIR")
		fmt.Println("6. BORRAR TODOS LOS DATOS")
		fmt.Print("INGRESA TU OPCION: ")

		// almacena numero en variable
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// creamos un NewReader llamado reader
			reader := bufio.NewReader(os.Stdin)

			// declaramos dos variables de tipo string para almacenar datos ingresados por el usuario
			var title, author string

			// Input prompt for title
			fmt.Print("Ingresar Titulo: ")
			// empezamos un for loop
			for {
				// leemos el input usando el reader y descartamos el valor de retorno de error usando _
				title, _ = reader.ReadString('\n')
				// borramos cualquier espacio al final o comienzo del string
				title = strings.TrimSpace(title)
				// si el titulo no es un string vacio salimos del for loop
				if title != "" {
					break
				}
			}

			// Input prompt for author
			fmt.Print("Ingresar Autor: ")
			for {
				author, _ = reader.ReadString('\n')
				author = strings.TrimSpace(author)
				if author != "" {
					break
				}
			}

			// Insert into database
			_, err = db.Exec("INSERT INTO users(title, author) VALUES (?, ?)", title, author)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error al agregar libro: %v\n", err)
			} else {
				fmt.Println("Libro Agregado Correctamente.")
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
					fmt.Println("Libro no encontrado.")
				} else {
					fmt.Println("Libro Borrado Exitosamente.")
				}
			}
		case 3:
			rows, err := db.Query("SELECT id, title, author FROM users")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error al mostrar usuarios: %v\n", err)
			} else {
				defer rows.Close()

				fmt.Println("*****************************************************************************")
				fmt.Println("ID\tTitle\t\t             Author")
				fmt.Println("_____________________________________________________________________________")
				for rows.Next() {
					var id int
					var title string
					var author string
					err = rows.Scan(&id, &title, &author)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error showing users: %v\n", err)
						break
					}
					output := fmt.Sprintf("%d   %-30s   %-30s", id, title, author)
					fmt.Println(output)
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
					row.AddCell().SetValue("Title")
					row.AddCell().SetValue("Author")

					// itera a través de las filas de resultados y agrega datos a las celdas
					for rows.Next() {
						var id int
						var title string
						var author string
						rows.Scan(&id, &title, &author)

						row := sheet.AddRow()
						row.AddCell().SetValue(id)
						row.AddCell().SetValue(title)
						row.AddCell().SetValue(author)
					}

					// guarda el archivo y muestra mensaje de éxito
					err = file.Save("bookstore.xlsx")
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error al guardar archivo: %v\n", err)
					} else {
						fmt.Println("Tabla de Usuarios Exportada a 'bookstore.xlsx'.")
					}
				}
			}
		case 5:
			fmt.Println("HASTA LUEGO!")
			return
		default:
			fmt.Println("Opcion no valida. Intente nuevamente.")

		case 6:
			result, err := db.Exec("DELETE FROM users")
			if err != nil {
				fmt.Fprintf(os.Stderr, "No se pudieron borrar datos: %v\n", err)
			} else {
				rowsAffected, _ := result.RowsAffected()
				if rowsAffected == 0 {
					fmt.Println("No hay datos para borrar")
				} else {
					fmt.Println("Datos borrados exitosamente")
				}
			} // fin case 6

		} // final switch statement

		fmt.Println("*****************************************************************************")

	} // final for loop principal

} // final funcion main
