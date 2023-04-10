package main

import (
	"fmt"
	"io/ioutil"

	
	"github.com/russross/blackfriday/v2"
)

func main() {
	mdFile := "example.md"
	htmlFile := "example.html"

	mdBytes, err := ioutil.ReadFile(mdFile)
	if err != nil {
		panic(err)
	}
	
	htmlBytes := blackfriday.Run(mdBytes)
	
	err = ioutil.WriteFile(htmlFile, htmlBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s convertido exitosamente a %s", mdFile, htmlFile)

}