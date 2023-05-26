package main

import (
	"crypto/rand"
	"fmt"
)

func main()  {
	var buf [4]byte // un slice de bytes de tamaño 4
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(err)
	}
	fmt.Println(buf) // output 4 numeros random entre 0 y 255
}