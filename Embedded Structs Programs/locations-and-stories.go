package main

import "fmt"

// struct comun en minuscula para que sea local a archivo
type quote struct {
	Text string
}
type Location struct {
	quote // embeded struct
	Name string
	Country string
}
type Story struct {
	quote // embeded struct
	Title string
	Author string
	Year int
}

func main() {
	 // creamos una locacion
	 l := Location {
		Name: "Prince Prospero's Castle",
		Country:"Italy" ,
		quote: quote{Text: "The Red Death had long devastated the country."},
	 }

	 // creamos una historia
	 h := Story {
		Title: "The Tell-Tale Heart",
		Author: "Edgar Allan Poe",
		Year: 1843,
		quote: quote{Text: "True! - nervous - very, very nervous I had been and am; but why will you say that I am mad?"},
	 }

	 // mostrar la informacion
	 fmt.Println("Titulo -", h.Title)
	 fmt.Println("Pais -", l.Country)
	 fmt.Println("Autor -",h.Author)
	 fmt.Println("Año -",h.Year)
	 fmt.Println("Locacion -",l.Name)
	 fmt.Printf("Quote Locacion:'%s'\n",l.Text)
	 fmt.Printf("Quote Libro:'%s'\n",h.Text)
}