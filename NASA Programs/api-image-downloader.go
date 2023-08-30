package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	imageURL := "https://api.nasa.gov/planetary/apod?api_key=gPLpgq6bmOWw9qzwEWaMZDDlSZG9Ndk73zx96TOG"
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Failed to fetch Image", err)
		return
	}

	defer resp.Body.Close()

	file, err := os.Create("nasa_img.jpg")
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Failed to save image", err)
		return
	}
	fmt.Println("NASA image downloaded and saved as nasa_image.jpg.")
}
