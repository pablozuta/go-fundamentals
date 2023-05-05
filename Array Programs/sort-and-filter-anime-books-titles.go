package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	titles := [] string {
	"Neuromancer",  
	"Snow Crash",  
	"Altered Carbon",  
	"Do Androids Dream of Electric Sheep?",  
	"The Diamond Age",  
	"Count Zero",  
	"Mona Lisa Overdrive",  
	"Virtual Light",  
	"Hardwired",  
	"Idoru",
	}

	// se ordenan alfabeticamente
	sort.Strings(titles)

	fmt.Println("Titulos ordenados alfabeticamente:")
    for _, title := range titles {
		fmt.Println(title)
	}

	// se dejan afuera los titulos que empiezan con 'N'
	filtered := make([]string, 0)
	for _, title := range titles {
		if !strings.HasPrefix(title, "N") {
			filtered = append(filtered, title)
		}
	}

	fmt.Println("Filtered titles:")
	for _, title := range filtered {
		fmt.Println(title)
	}
}