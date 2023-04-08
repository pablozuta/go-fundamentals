package main  
  
import (  
	"bufio"  
	"fmt"  
	"github.com/jung-kurt/gofpdf"  
	"os"  
) 

func main() {
	// abrir el archivo txt
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("no se pudo leer el archivo")
	}

	defer inputFile.Close()

	// crear un archivo pdf
	pdf := gofpdf.New("P", "mm", "A4", "")
	// fuente y tamaño de fuente
	pdf.SetFont("courier", "", 10)

	// crea una pagina nueva 
	pdf.AddPage() 
	  
	// lee cada linea del archivo txt y lo escribe al PDF 
	scanner := bufio.NewScanner(inputFile)  
	y := 10 // posicion de inicio  
	
	for scanner.Scan() {  
	text := scanner.Text()  
	pdf.Text(10, float64(y), text)  
	y += 7 // baja por la coordenada y para crear una siguiente linea 
	}  
	if err := scanner.Err(); err != nil {  
	panic(err)  
	}  
	  
	// graba el archivo PDF  
	err = pdf.OutputFileAndClose("output.pdf")  
	if err != nil {  
	panic(err)  
	}  
	  
	fmt.Println("PDF generado exitosamente.") 
}