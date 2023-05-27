package main

import "fmt"

// definimos un interface que representa un catto
type Cat interface {
	Name() string
	Sound() string
}

// definimos un struct para un gato domestico
type DomesticCat struct {
	NameVal string
}

// implementamos el metodo Name para el struct gato domestico
func (d DomesticCat) Name() string {
	return d.NameVal
}

// implementamos el metodo Sound para el struct gato domestico
func (d DomesticCat) Sound() string {
	return "Meow"
}

// definimos un struct para un leon
type Lion struct {
	NameVal string
}

// implementamos el metodo Name para el struct leon
func (l Lion) Name() string {
	return l.NameVal
}

// implementamos el metodo Sound para el struct leon
func (l Lion) Sound() string {
	return "Roar"
}

func main() {
	// creamos un slice de gatos
	cats := []Cat{
		DomesticCat{NameVal: "Mermelon"},
		Lion{NameVal: "Simba"},
		DomesticCat{NameVal: "Lorena"},
	}

	// mostramos los nombres y sonidos del slice
	for _, cat := range cats {
		fmt.Printf("%s says %s\n", cat.Name(), cat.Sound())
	}
}
