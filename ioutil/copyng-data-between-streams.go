
package main  
  
import (  
	"fmt"  
	"io"  
	"net/http"  
	"os"  
)  
  
func main() {  
	// Downloading a file and saving it locally  
	response, err := http.Get("https://w7.pngwing.com/pngs/802/825/png-transparent-redbubble-polite-cat-meme-funny-cat-meme-thumbnail.png")  
	if err != nil {  
	fmt.Println("Error downloading file:", err)  
	return  
	}  
	defer response.Body.Close()  
	  
	file, err := os.Create("image.png")  
	if err != nil {  
		fmt.Println("Error creating file:", err)  
		return  
	}  
	defer file.Close()  
	  
	_, err = io.Copy(file, response.Body)  
	if err != nil {  
		fmt.Println("Error copying data:", err)  
	return  
	}  
	fmt.Println("File downloaded and saved as image.png")  
}  
