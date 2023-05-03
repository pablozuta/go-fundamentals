package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    // abre el archivo sqlite
    db, err := sql.Open("sqlite3", "bookstore.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() // espera hasta que se ejecute main para cerrar la base de datos

    // selecciona todo desde la tabla users
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // se crea un objeto CSV writer que recibira los datos
    file, err := os.Create("output.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // escribe el row de headers 
    columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }
    writer.Write(columns)

    // hace una iteracion por las filas y las escribe en el archivo CSV
    values := make([]sql.NullString, len(columns))
    valuePtrs := make([]interface{}, len(columns))
    for i := range columns {
        valuePtrs[i] = &values[i]
    }
    for rows.Next() {
        if err := rows.Scan(valuePtrs...); err != nil {
            log.Fatal(err)
        }
        record := make([]string, len(values))
        for i, value := range values {
            if value.Valid {
                record[i] = value.String
            }
        }
        writer.Write(record)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    } else {
		fmt.Println("Archivo CSV Creado Exitosamente")
	}
}
