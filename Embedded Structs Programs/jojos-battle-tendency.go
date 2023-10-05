// Here's a program that uses embedded structs and interfaces to represent characters from the anime "Jojo's Bizarre Adventure: Battle Tendency", which is the second season of the series.

// This program defines multiple levels of embedding and interface implementation. The JosephJoestar struct is embedded in the CaesarZeppeli struct, which implements the HamonUser interface. The Wamuu struct is embedded in the Kars struct, which implements the Vampire interface. This program also demonstrates how to use type assertions to determine if a given character implements a certain interface.

package main

import "fmt"

// Character is an interface that defines methods for interacting with a character
type CharacterI interface {
	GetName() string
	GetAge() int
	GetAbilities() []string
}

// Ability is a struct that represent a character's ability
type Ability struct {
	Name        string
	Description string
}

//JosephJoestar is a struct that represents the character Joseph Joestar
type JosephJoestar struct {
	Name     string
	Age      int
	English  string
	Japanese string
	Ability  Ability // embeded Ability struct (Fields: Name, Description)
}

// the GetName method returns the name of the character
func (c JosephJoestar) GetName() string {
	return c.Name
}

// the GetAge method returns the name of the character
func (c JosephJoestar) GetAge() int {
	return c.Age
}

// GetAbilities return a list of abilities of the character
func (c JosephJoestar) GetAbilities() []string {
	return []string{c.Ability.Name, c.Ability.Description}
}

// HamonUser is an interface that defines methods for interacting with a Hamon user
type HamonUser interface {
	CharacterI
	GetRippleTechniques() []string
}

// CaesarZeppeli is a struct that represent the character Caesar Zeppeli
type CaesarZeppeli struct {
	JosephJoestar // embed the JosephJoestar struct
	BirthYear     int
	Techniques    []string
}

// GetRipple techniques return the techniques of the character
func (c CaesarZeppeli) GetRippleTechniques() []string {
	return c.Techniques
}

// Wamuu is a struct that represent the character Wamuu
type Wamuu struct {
	Name      string
	Age       int
	Abilities []Ability // embeded slice of Ability struct
}

// the GetName method returns the name of the character
func (c Wamuu) GetName() string {
	return c.Name
}

// the GetAge method returns the name of the character
func (c Wamuu) GetAge() int {
	return c.Age
}

// GetAbilities returns the list of abilities of the Wamuu character
func (c Wamuu) GetAbilities() []string {
	abilities := []string{}
	for _, a := range c.Abilities {
		abilities = append(abilities, a.Name, a.Description)
	}
	return abilities
}

// Vampire is an interface that defines methods for interacting with a vampire
type Vampire interface {
	CharacterI
	GetWeakness() string
}

// Kars is a struct that represent the character Kars
type Kars struct {
	Wamuu    // Embed the Wamuu struct
	Weakness string
}

// GetWeakness return the weakness of the character
func (c Kars) GetWeakness() string {
	return c.Weakness
}

func main() {
	// create some characters
	joseph := JosephJoestar{
		Name:     "Joseph Joestar",
		Age:      18,
		English:  "Joseph Joestar",
		Japanese: "ジョセフ・ジョースター",
		Ability: Ability{
			Name:        "Hamon",
			Description: "A breathing technique that uses the power of the sun.",
		},
	}

	caesar := CaesarZeppeli{
		JosephJoestar: joseph, // embed the Joseph Joestar struct
		BirthYear:     1956,
		Techniques: []string{
			"Bubble Cutter",
			"Bubble Barrier",
			"Bubble Launcher",
		},
	}

	wamuu := Wamuu{
		Name: "Wamuu",
		Age:  150,
		Abilities: []Ability{
			{Name: "Wind Mode", Description: "Allows the user to control wind."},
			{Name: "Santana's Arm Blade", Description: "Blades that extends from the user's arms"},
		},
	}

	kars := Kars{
		Wamuu:    wamuu, // embed the Wamuu struct
		Weakness: "SunLight",
	}

	// define a list of characters
	characters := []CharacterI{joseph, caesar, wamuu, kars}

	// loop through the characters and print their information
	for _, c := range characters {
		fmt.Printf("Name: %s\n", c.GetName())
		fmt.Printf("Age: %d\n", c.GetAge())

		// invocamos el metodo GetAbilities y recibimos un []string
		abilities := c.GetAbilities() 
		for _, a := range abilities {
			fmt.Println("-", a)
		}
		fmt.Printf("Abilities: %v\n", c.GetAbilities())

		// If the character is a Hanon user print their ripple techniques
		if hu, ok := c.(HamonUser); ok {
			fmt.Printf("Ripple Techniques: %v\n", hu.GetRippleTechniques())
		}

		// If the character is a vampire print their weakness
		if v, ok := c.(Vampire); ok {
			fmt.Printf("Weakness: %s\n", v.GetWeakness())
		}

		fmt.Println()

	}
}
