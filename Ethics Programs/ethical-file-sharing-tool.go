// This program, "Ethical File Sharing Tool," allows users to upload files with ethical considerations and promotes responsible file sharing practices.
package main

import "fmt"

type File struct {
	Name     string
	Size     int
	Ethical  bool
	Location string
}

type FileSharingApp struct {
	Files []File
}

func main() {
	app := FileSharingApp{}
	fmt.Println("Welcome to the file sharing tool")

	for {
		fmt.Print("Enter the file name(or 'exit' to finish): ")
		var name string
		fmt.Scan(&name)

		if name == "exit" {
			break
		}

		fmt.Print("Enter the file size (in MB): ")
		var size int
		fmt.Scan(&size)

		fmt.Print("Is this file ethical?: ")
		var ethical bool
		fmt.Scan(&ethical)

		fmt.Print("Enter the file location: ")
		var location string
		fmt.Scan(&location)

		file := File{Name: name, Size: size, Ethical: ethical, Location: location}
		app.AddFile(file)
	}

	fmt.Println("Files Uploaded:")
	for _, file := range app.Files {
		fmt.Printf("Name: %s, Size: %dMB, Ethical: %v, Location: %s\n", file.Name, file.Size, file.Ethical, file.Location)
	}
}

func (app *FileSharingApp) AddFile(file File) {
	app.Files = append(app.Files, file)
}
