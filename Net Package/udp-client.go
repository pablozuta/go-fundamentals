package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		fmt.Println("No se pudo conectar", err)
		return
	}
	defer conn.Close()

	message :=[]byte ("Hello Server")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("No se pudo enviar el mensaje", err)
		return
	}
	fmt.Println("Message Sent:", string(message))
}