package main

import "fmt"

// struct para representar a personaje
type Character struct {
	Name      string
	Age       int
	Abilities []string
}

// struct para representar el anime
type Akira struct {
	Title      string
	Characters []Character
}

func main() {
	// crear un slice de Character structs
	characters := []Character{
		Character{
			Name: "Kaneda",
			Age:  16,
			Abilities: []string{
				"Motorcycle Riding", "Hand on hand combat"},
		},
		Character{
			Name: "Tetsuo",
			Age:  15,
			Abilities: []string{
				"Telekinesis",
				"Superhuman Strength",
				"Pyrokinesis",
			},
		},
		Character{
			Name: "Key",
			Age:  17,
			Abilities: []string{
				"Physic Powers", "Espionage",
			},
		},
	}
	// creamos struct akira y le asignamos el slice de los personajes que creamos
	akira := Akira{
		Title:      "Akira",
		Characters: characters,
	}
	// mostrar los personajes del anime 
	fmt.Println("Personajes del Anime Akira:")
	fmt.Println()
	for _, character := range akira.Characters {
		fmt.Printf("Nombre:%s (edad %d)\n", character.Name, character.Age)
		fmt.Println(" Habilidades:")
		for _, ability := range character.Abilities {
			fmt.Printf(" - %s\n", ability)
		}
	}

}
