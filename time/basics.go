package main

import (
	"fmt"
	"time"
)

func main() {
	// obtener el tiempo actual
	ahora := time.Now()
	fmt.Println("El tiempo actual es:", ahora)

	// sleep por 1 segundo
	time.Sleep(1 * time.Second)

	// obtener el tiempo pasado desde el ultimo sleep
	elapsed := time.Since(ahora)
	fmt.Println("Tiempo transcurrido:", elapsed)

	// parse time string
	layout := "2006-01-02T15:04:05.000Z"
	value := "2023-04-01T12:00:00.000Z"
	parsedTime, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println("No se pudo hacer parsing de time , error:", err)
		return
	}
	fmt.Println("Parsed Time:", parsedTime)

}
