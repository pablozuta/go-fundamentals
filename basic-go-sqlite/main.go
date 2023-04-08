package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	
)

func main() {
	// abre una coneccion con la base de datos
	db, err := sql.Open("sqlite3", "records.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// crea una tabla si no existe
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		panic(err)
	}

	fmt.Println("se creo tabla 'users' en base de datos 'records' exitosamente")
}
