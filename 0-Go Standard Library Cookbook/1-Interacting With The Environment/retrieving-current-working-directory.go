// Another useful source of information for the application is the directory, where the program binary is located. With this information, the program can access the assets and files collocated with the binary file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// The Executable function returns the absolute path of
	// the binary that is executed
	// (unless the error is returned).
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// path to executable file
	fmt.Println(ex)

	// resolve the directory of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable Path:" + exPath)

	// use EvalSymlinks to get the real path.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Symlink evaluated:" + realPath)
}
