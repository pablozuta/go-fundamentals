package main

import "fmt"

func main() {
	day := '日'
	fmt.Println(day) // output 26085
	fmt.Println(string(day)) // output 日
	fmt.Printf("%c\n", day) // output 日

	saludo := []rune {'h', 'o', 'l', 'a'}
	fmt.Println(len(saludo)) // output 4
	
	for _, item := range saludo {
		fmt.Println(string(item))
	}
}