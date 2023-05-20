package main

import "fmt"

type Model struct {
	Name    string
	Version string
	Weight  float64
}
type Robot struct {
	Model          // embed the model struct
	Name           string
	SerialNumber   string
	ProductionYear int
}
type Manufacturer struct {
	Model             // embed the model struct
	Name              string
	ProductionCountry string
}

func main() {
	// creamos un modelo
	m := Model{
		Name:    "Tachikoma",
		Version: "AI Robot",
		Weight:  120.5,
	}
	// creamos un robot
	r := Robot{
		Model:          m,
		Name:           "Tachi",
		SerialNumber:   "003",
		ProductionYear: 2030,
	}
	// creamos un manufacturador
	mf := Manufacturer{
		Model:             m,
		Name:              "Tachikoma Factory",
		ProductionCountry: "Japan",
	}

	// mostramos la data del robot
	fmt.Println("Nombre                ",r.Name)
	fmt.Println("Numero de Serie       ",r.SerialNumber)
	fmt.Println("Año de Produccion     ",r.ProductionYear)
	fmt.Println("Modelo                ",r.Model.Name)
	fmt.Println("Version               ",r.Model.Version)
	fmt.Println("Peso                  ",r.Model.Weight)
	fmt.Println("Manufacturer          ",mf.Name)
	fmt.Println("Manufacturer Country  ",mf.ProductionCountry)


}
