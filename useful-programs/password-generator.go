package main

import (
	"fmt"
	"math/rand"
	"time"
)

// creamos una funcion que acepta como argumento un integer que sera la longitud de la contraseña y retorna un string con la contraseña generada
func generarContraseña(length int) string {
	var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()" 
	var contraseña = make([]byte, length)
	rand.Seed(time.Now().UnixNano())

	for i:= 0; i < length; i++ {
		contraseña[i] = chars[rand.Intn(len(chars))]

	}
	return string(contraseña)
}

func main() {
	contraseña := generarContraseña(18) // variable que guardara el string obtenido
	fmt.Printf("la contraseña es: %s", contraseña)
}