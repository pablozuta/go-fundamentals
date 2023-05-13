package main

import "fmt"

// la funcion recibe un slice de strings 'A'
func commonChars(A []string) []string {
	// creamos un 2D array de la longitud del array que recibamos para guardar los resultados
	counts := make([][26]int, len(A))

	for i, s := range A {
		for _, c := range s {
			counts[i][c-'a']++
		}
	}
	// Find the common characters
	result := []string{}
	for i := 0; i < 26; i++ {
		minCount := counts[0][i]
		for j := 1; j < len(counts); j++ {
			if counts[j][i] < minCount {
				minCount = counts[j][i]
			}
		}
		for k := 0; k < minCount; k++ {
			result = append(result, string(i+'a'))
		}
	}

	return result
}

func main() {
	A := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(A))
}
