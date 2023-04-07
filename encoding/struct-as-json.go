package main

import (
	"encoding/json"
	"fmt"
)

// se crea struct
type Person struct {
	Name string
	Age int
}

func main() {
	persona := Person {"Alice", 24} // crea persona usando struct
	personaJSON, err := json.Marshal(persona) // hace encoding de struct a JSON guardandolo en variable

	if err != nil {
		fmt.Println("No se pudo hacer encoding")
		return
	}
	
	fmt.Println("go struct:", persona)
	fmt.Println("JSON object:" ,string(personaJSON))
}