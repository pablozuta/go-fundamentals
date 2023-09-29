// sample program to show how a bytes.Buffer can also be used with the io.Copy function
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	// write a string to the buffer
	b.Write([]byte("Hola Golang"))

	// use Fprintf to concatenate a string to the buffer
	fmt.Fprintf(&b, " and Gopher!")

	// write the content of the buffer to Stdout
	io.Copy(os.Stdout, &b)
}
