package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, "Writer From Bufio")
	writer.Flush()
}