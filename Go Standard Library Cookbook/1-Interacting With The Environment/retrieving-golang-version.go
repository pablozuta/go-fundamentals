package main

import (
	"fmt"
	"runtime"
)

const info = `
Appication "%s" starting.
The binary was build by GO: %s 
`

func main() {
	fmt.Printf(info, "retrieving golang version", runtime.Version())
}
