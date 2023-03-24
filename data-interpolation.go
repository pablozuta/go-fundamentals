package main

import "fmt"

func main() {
	// In Go, you can perform string interpolation using the fmt package
	// which provides a variety of formatting options for strings.
	name := "Harold Incandenza"
	age := 18
	fmt.Printf(`You are %s , %d , date of secondary-school graduation approximately one
    month from now, attending the Enfield Tennis Academy, Enfield, Massachusetts, a boarding school, where you reside.\n`, name, age)
	fmt.Println(" ")		

	// You can also use the Sprintf() function to create
	// a formatted string that you can later print or use in other ways
	formattedString := fmt.Sprintf(`You are %s , %d , date of secondary-school graduation approximately one month from now, attending the Enfield Tennis Academy, Enfield, Massachusetts, a boarding school, where you reside.\n`, name, age)
	fmt.Println(formattedString)
	fmt.Println(" ")

	// FLOATS
	temperature := 23.5
	// formatea la data usando dos decimales
	fmt.Printf("The temperature outside is %.2f degrees Celsius\n", temperature)
	fmt.Println(" ")

	// BOOLEANS
	isRaining := false
	fmt.Printf("Is Raining? %t \n", isRaining)
	fmt.Println(" ")

	// STRINGS
	var firstName = "Don"
	var lastName = "DeLillo"
	fmt.Printf(`%s %s is an American novelist, playwright, and essayist who was born on November 20, 1936, in New York City. He is widely regarded as one of the most important and influential writers of his generation `, firstName, lastName)
	fmt.Println(" ")
	fmt.Println(" ")

	// POINTERS
	var edad = 40
	var edadPointer = &edad
	fmt.Printf("My age is %d (stored at %p)", *edadPointer, edadPointer)

}
