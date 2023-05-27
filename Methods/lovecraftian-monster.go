package main

import "fmt"

// creamos un struct con dos fields de tipo string y de tipo map
type Monster struct {
	Type       string
	Attributes map[string]string
}

// creamos un metodo para el struct Monster
func (m Monster) Description() string {
	// el valor string se usa al medio de una frase que retornaremos
	description := fmt.Sprintf("This is a %s with the following attributes:\n", m.Type)
	// hacemos un for loop por el mapa , usando el valor key y el valor value
	for attribute, value := range m.Attributes {
		// concatenamos la frase que ya teniamos con el valor de cada value pair del mapa
		description += fmt.Sprintf("- %s: %s\n", attribute, value)
	}
	return description
}

func main() {
	shoggoth := Monster{Type: "Shoggoth", Attributes: map[string]string{"Size": "Enormous", "Shape": "Amorphous", "Abilities": "Shape-shifting, Telepathy"}}
	deepOne := Monster{Type: "Deep One", Attributes: map[string]string{"Appearance": "Fish-like, Humanoid", "Culture": "Worships Cthulhu", "Habitat": "Underwater Cities"}}

	fmt.Println(shoggoth.Description())
	fmt.Println(deepOne.Description())
}
