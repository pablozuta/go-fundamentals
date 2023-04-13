package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "[contraseña base de datos]"
	dbname   = "[nombre base de datos a crear]"
)

func main() {
	// conexion con postgres
	// sslmode=disable is an option that disables SSL encryption for the connection
	// (not recommended for production)
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "postgres")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(connStr)

	// se crea la base de datos
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		panic(err)
	}

	// conneccion a la base de datos
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// creamos una tabla
	_, err = db.Exec(`CREATE TABLE users (  
             id SERIAL PRIMARY KEY,  
             username TEXT,  
             email TEXT,  
             address TEXT  
        )`)
	if err != nil {
		panic(err)
	}

	fmt.Println("base de datos creada exitosamente!")
}
