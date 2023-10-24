// This program is an "Ethical Decision Support Tool for Online Shopping" that allows users to enter a product name and receive an ethical rating for that product. It encourages ethical consumption and helps users make informed choices about the products they buy.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Ethical Decision Support Tool for Online Shopping!")
	fmt.Print("Enter the name of a product: ")

	var nameProduct string
	fmt.Scan(&nameProduct)

	ethicalRating := rateProductEthics(nameProduct)
	fmt.Printf("Ethical Rating for %s: %s\n", nameProduct, ethicalRating)
}

func rateProductEthics(productName string) string {
	// In a real-world scenario, this function would use external data and algorithms to rate the ethics of the product.
	// For simplicity, let's assume it always returns a positive rating.
	return "Positive (Ethical product)"
}
