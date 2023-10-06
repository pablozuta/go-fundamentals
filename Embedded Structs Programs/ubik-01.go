// In this program, there are four struct types: UbikCharacter, UbikAbility, UbikPsiAbility, and UbikCharacterWithPsi. The UbikPsiAbility struct embeds the UbikAbility struct using the syntax UbikAbility. The UbikCharacterWithPsi struct embeds both the UbikCharacter and UbikPsiAbility structs using the syntax UbikCharacter and UbikPsiAbility.

// In the main function, a new UbikCharacterWithPsi instance is created with a name of "Joan Reiss", an age of 27, a gender of "Female", an ability name of "Precognition", an ability description of "Can see the future", and an ability level of 3. The program then prints out Joan's information and her ability information using fmt.Printf.
package main

import "fmt"

type UbikCharacter struct {
	Name   string
	Age    int
	Gender string
}

type UbikAbility struct {
	Name        string
	Description string
}

type UbikPsiAbility struct {
	UbikAbility
	Level int
}

type UbikCharacterWithPsi struct {
	UbikCharacter
	UbikPsiAbility
}

func main() {
	// create a new UbikCharacterWithPsi instance
	joan := UbikCharacterWithPsi{
		UbikCharacter: UbikCharacter{Name: "Joan Reiss", Age: 27, Gender: "Female"},
		UbikPsiAbility: UbikPsiAbility{
			UbikAbility: UbikAbility{Name: "Precognition", Description: "Can see the Future"},
			Level:       3,
		},
	}

	// Print out Joans Information
	fmt.Printf("Character Name: %s\n", joan.Name)
	fmt.Printf("Age: %d\n", joan.Age)
	fmt.Printf("Gender: %s\n", joan.Gender)
	fmt.Printf("Ability Name: %s\n", joan.UbikAbility.Name)
	fmt.Printf("Ability Description: %s\n", joan.Description)
	fmt.Printf("Ability Level: %d\n", joan.Level)
}
