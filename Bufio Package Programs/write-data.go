package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var frase = "El Aleph"
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(writer, "Writer From Bufio")
	fmt.Fprintln(writer, "Writer From Bufio Dos")
	fmt.Fprint(writer, frase )
	writer.Flush()
}