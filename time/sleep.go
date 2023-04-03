package main

import (
	"fmt"
	"time"
)

func main() {
	duracion := 10 * time.Second
	fmt.Println("Sleep por: ", duracion)

	time.Sleep(duracion)
	fmt.Println("DESPERTAR!")
}