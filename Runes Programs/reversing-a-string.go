package main

import "fmt"

func reversingString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "Hello, 世界"
	reversed := reversingString(str)
	fmt.Println("Frase original:", str)
	fmt.Println("Frase en reversa:", reversed)

}
