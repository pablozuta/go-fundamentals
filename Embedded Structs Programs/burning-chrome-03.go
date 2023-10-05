// In this program, there are two struct types: fiction and shortStory. The shortStory struct embeds the fiction struct using the syntax fiction. This means that the fields of the fiction struct can be accessed directly from the shortStory struct.

// In the main function, a new shortStory instance is created with the title "Burning Chrome", the author "William Gibson", the genre "Science Fiction", and a word count of 12000. The program then prints out the short story's information using fmt.Printf.
package main

import "fmt"

type fiction struct {
	title  string
	author string
	genre  string
}

type shortStory struct {
	fiction
	wordCount int
}

func main() {
	// create a new shortStory instance
	burningChrome := shortStory{
		fiction: fiction{
			title:  "Burning Chrome",
			author: "William Gibson",
			genre:  "Cyber-Punk",
		},
		wordCount: 12000,
	}
	// print out the shortHistory information
	fmt.Printf("Title: %s\n", burningChrome.title)
	fmt.Printf("Author: %s\n", burningChrome.author)
	fmt.Printf("Genre: %s\n", burningChrome.genre)
	fmt.Printf("Word Count: %d\n", burningChrome.wordCount)
}
