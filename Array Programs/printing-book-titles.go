// This program takes an array of cyberpunk book titles as input and displays them one by one.  
package main

import "fmt"

func main()  {
	titles := [] string {"Neuromancer", "Snow Crash", "Altered Carbon", "Do Android Dream of Electric Sheep", "The Diamond Age", "Count Zero", "Mona Lisa Overdrive", "Virtual Light", "Hardwired", "Idoru"}
	fmt.Println("LIST OF CYBERPUNK BOOKS:")

	for _, title := range titles {
		fmt.Println("-" + title)
	}
}