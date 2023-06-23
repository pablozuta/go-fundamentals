// go run imageresizer.go input.jpg 0.5 
package main  
  
import (  
    "fmt"  
    "image"  
    "image/jpeg"  
    "image/png"  
    "os"  
    "path/filepath"  
    "strconv"  
    "strings"  
)  
  
func main() {  
     args := os.Args  

    // Open the input file for reading  
    inputFile, err := os.Open(args[1])  
    if err != nil {  
        panic(err)  
    }  
    defer inputFile.Close()  

    // Determine the output file format based on the file extension of the input file  
    var outputFileFormat string  
    if strings.ToLower(filepath.Ext(args[1])) == ".jpg" || strings.ToLower(filepath.Ext(args[1])) == ".jpeg" {  
    outputFileFormat = "jpg"  
    } else if strings.ToLower(filepath.Ext(args[1])) == ".png" {  
    outputFileFormat = "png"  
    } else {  
    panic("Unsupported file format")  
    }  

    // Decode the input file into an image.Image object  
    img, _, err := image.Decode(inputFile)  
    if err != nil {  
    panic(err)  
    }  

    // Parse the scaling factor  
    scaleFactor, err := strconv.ParseFloat(args[2], 64)  
    if err != nil {  
    panic(err)  
    }  

    // Calculate the dimensions of the resized image  
    resizedWidth := int(float64(img.Bounds().Dx()) * scaleFactor)  
    resizedHeight := int(float64(img.Bounds().Dy()) * scaleFactor)  

    // Create a new image with the calculated dimensions  
    resizedImg := image.NewRGBA(image.Rect(0, 0, resizedWidth, resizedHeight))  

    // Scale the original image into the new image  
    for x := 0; x < resizedWidth; x++ {  
    for y := 0; y < resizedHeight; y++ {  
    resizedImg.Set(x, y, img.At(int(float64(x)/scaleFactor), int(float64(y)/scaleFactor)))  
    }  
    }  

    // Create the output file and encode the resized image into it  
    outputFile, err := os.Create(fmt.Sprintf("resized.%s", outputFileFormat))  
    if err != nil {  
        panic(err)  
    }  
    defer outputFile.Close()  

    if outputFileFormat == "jpg" {  
        jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 80})  
    } else {  
        png.Encode(outputFile, resizedImg)  
    } 
	fmt.Println("Archivo generado exitosamente!") 
}