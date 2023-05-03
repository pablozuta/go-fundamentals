package main

import (
	"fmt"
	"net"
)


func main()  {
	conexion, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error creando server", err)
		return 
	}

	defer conexion.Close()

	mensaje := "Hello server."
	_, err = conexion.Write([]byte(mensaje))
	if err != nil {
		fmt.Println("Error enviando mensaje", err)
		return 
	}
	fmt.Println("Mensaje enviado ", mensaje)
}