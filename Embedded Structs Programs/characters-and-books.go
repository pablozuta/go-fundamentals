package main

import "fmt"

// tanto el personaje como el libro usaran este struct
type Quote struct {
	Text string
}

type Character struct {
	Quote // Embed the quote struct
	Nombre string
	Edad int
}
type Book struct {
	Quote // embed the quote struct
	Titulo string
	Autor string
	Año int
}
func main() {
	// creamos un personaje (quote, name and age)
	c := Character {
		Edad: 50, // la data puede estar en cualquier orden 
		Nombre: "Montresor",  
		Quote: Quote{Text: "The thousand injuries of Fortunato I had borne as I best could, but when he ventured upon insult, I vowed revenge."},  
	}
	// Creamos un libro (quote, titulo, autor , año)
	b := Book{  
		Quote: Quote{Text: "During the whole of a dull,  dark, and soundless day , in the autumn of the year, when the clouds  hung oppressively low  in the heavens, I had been  passing alone, on horseback,  through a singularly dreary tract of country; and at length found myself, as the shades of the evening drew on, within view of the melancholy House of Usher."},  
		Titulo: "The Fall of the House of Usher",  
		Autor: "Edgar Allan Poe",  
		Año: 1839,  
		}  
		  
		// mostramos informacion sobre el personaje
		fmt.Println("***********DATOS DEL PERSONAJE****************")  
		fmt.Println("Nombre:", c.Nombre) // nombre personaje  
		fmt.Println("Edad:", c.Edad)  // edad del personaje
		fmt.Printf("Quote:'%s'\n", c.Text)  // quote del personaje
		fmt.Println()
		fmt.Println()
		
		// Print some information about the book  
		fmt.Println("*************DATOS DEL LIBRO*****************")  
		fmt.Println("Titulo:", b.Titulo)  // titulo del libro
		fmt.Println("Autor:", b.Autor)  // autor del libro
		fmt.Println("Año:", b.Año)  // año
		fmt.Printf("Quote:'%s'", b.Text) // frase del libro
		fmt.Println()
}