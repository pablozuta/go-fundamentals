package main

import (
	"bytes"

	"io"
	"os"
)

func main() {
	// creamos un byte slice
	data := []byte{0x48, 0x65, 0x64}

	// creamos un new reader que leera desde el byte slice
	reader := bytes.NewReader(data)

	// copiar la data del reader a stdout
	io.Copy(os.Stdout, reader)

}
