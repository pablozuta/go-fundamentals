// This program, the "Ethical Food Consumption Tracker," allows users to track their food consumption and promotes ethical food choices by marking items as ethical or not.
package main

import (
	"fmt"
	"strings"
)

type FoodItem struct {
	Name        string
	Ethical     bool
	Calories    int
	Ingredients []string
}

type FoodTracker struct {
	FoodItems []FoodItem
}

func main() {
	tracker := FoodTracker{}
	fmt.Println("Welcome to the Ethical Food Consumption Tracker!")

	for {
		fmt.Print("Enter the name of a food item (or 'exit' to finish): ")
		var name string
		fmt.Scan(&name)
		if name == "exit" {
			break
		}

		fmt.Print("Is this food item ethical? (true/false): ")
		var ethical bool
		fmt.Scan(&ethical)

		fmt.Print("Enter the number of calories: ")
		var calories int
		fmt.Scan(&calories)

		fmt.Println("Enter the ingredients (comma-separated, e.g., ingredient1,ingredient2): ")
		var ingredientsInput string
		fmt.Scan(&ingredientsInput)
		ingredients := strings.Split(ingredientsInput, ",")

		foodItem := FoodItem{Name: name, Ethical: ethical, Calories: calories, Ingredients: ingredients}
		tracker.AddFoodItem(foodItem)
	}

	fmt.Println("Food items tracked:")
	for _, item := range tracker.FoodItems {
		fmt.Printf("Name: %s, Ethical: %v, Calories: %d, Ingredients: %v\n", item.Name, item.Ethical, item.Calories, item.Ingredients)
	}
}

func (tracker *FoodTracker) AddFoodItem(foodItem FoodItem) {
	tracker.FoodItems = append(tracker.FoodItems, foodItem)
}
