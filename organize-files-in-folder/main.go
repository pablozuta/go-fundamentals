package main

import (
	"fmt"
	"os"
	
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run main.go [path-hacia-directorio]")
		return
	}

	folderPath := os.Args[1]

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		extension := filepath.Ext(path)
		newFolderPath := filepath.Join(folderPath, extension[1:])
		err = os.MkdirAll(newFolderPath, 0755)
		if err != nil {
			return err
		}
		
		newFilePath := filepath.Join(newFolderPath, info.Name())
		err = os.Rename(path, newFilePath)
		if err != nil {
			return err
		}
		return nil
		
	})
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("archivos organizados exitosamente")
}