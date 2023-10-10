// In this program, there are four struct types: `UbikObject`, `UbikBuilding`, `UbikProduct`, and `UbikStore`. The `UbikBuilding` struct embeds the `UbikObject` struct using the syntax `UbikObject`. The `UbikProduct` struct also embeds the `UbikObject` struct using the syntax `UbikObject`. The `UbikStore` struct embeds both the `UbikBuilding` and a slice of `UbikProduct` structs using the syntax `UbikBuilding` and `[]UbikProduct`.

// In the `main` function, a new `UbikStore` instance is created with a name of "Runciters Store", an age of 50, a company of "Runciter Associates", a building location of "San Francisco", and a slice of three `UbikProduct` instances. The program then prints out the Runciters Store information and the product information for each product in the store using `fmt.Printf` and a `for` loop.
package main

import "fmt"

type UbikObject struct {
	Name    string
	Age     int
	Company string
}

type UbikBuilding struct {
	UbikObject
	Location string
}

type UbikProduct struct {
	UbikObject
	Price int
}

type UbikStore struct {
	UbikBuilding
	Products []UbikProduct
}

func main() {
	// Create a new UbikStore instance
	runcitersStore := UbikStore{
		UbikBuilding: UbikBuilding{
			UbikObject: UbikObject{
				Name:    "Runciters Store",
				Age:     50,
				Company: "Runciter Associates",
			},
			Location: "San Francisco",
		},
		Products: []UbikProduct{
			UbikProduct{
				UbikObject: UbikObject{
					Name:    "Joe Chip's Cigarettes",
					Age:     2,
					Company: "Runciter Associates",
				},
				Price: 10,
			},
			UbikProduct{
				UbikObject: UbikObject{
					Name:    "Moratorium Beer",
					Age:     3,
					Company: "Beloved Brethren Moratorium",
				},
				Price: 5,
			},
			UbikProduct{
				UbikObject: UbikObject{
					Name:    "Hollis's Womb-Filling Station",
					Age:     1,
					Company: "Hollis International",
				},
				Price: 20,
			},
		},
	}

	// Print out the Runciters Store information
	fmt.Printf("Store Name: %s\n", runcitersStore.Name)
	fmt.Printf("Age: %d\n", runcitersStore.Age)
	fmt.Printf("Company: %s\n", runcitersStore.Company)
	fmt.Printf("Building Location: %s\n", runcitersStore.Location)

	// Print out the Runciters Store products information
	fmt.Println("Products:")
	for _, product := range runcitersStore.Products {
		fmt.Printf("Product Name: %s\n", product.Name)
		fmt.Printf("Age: %d\n", product.Age)
		fmt.Printf("Company: %s\n", product.Company)
		fmt.Printf("Price: %d\n", product.Price)
		fmt.Println()
	}
}
