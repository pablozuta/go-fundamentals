package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	
)

func main() {
    // variables
	var nombreDB = "users.db"
	var nombreTabla = "users"

	// abre una coneccion con la base de datos
	db, err := sql.Open("sqlite3", nombreDB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// crea una tabla si no existe
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		panic(err)
	}


	fmt.Printf("se creo tabla %s en base de datos %s exitosamente", nombreTabla, nombreDB)
}
