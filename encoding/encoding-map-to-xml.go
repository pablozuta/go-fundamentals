package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	m := map[string]int{"A Love Supreme": 1, "Kind of Blue": 2}
	b, err := xml.Marshal(m)

	if err != nil {
		fmt.Println("no se pudo realizar encoding")
	}
	fmt.Println(string(b))
}