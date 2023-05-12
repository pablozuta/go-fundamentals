package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)


func main() {
	data := []byte ("Burning Chome")
	var compressed bytes.Buffer
	
	gzWriter := gzip.NewWriter(&compressed)
	_, err := gzWriter.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = gzWriter.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compressed Data")	
	fmt.Printf("%v\n", compressed.Bytes())	
	
	gzReader, err := gzip.NewReader(&compressed)
	if err != nil {
		fmt.Println(err)
		return
	}
	var decompressed bytes.Buffer
	_, err = decompressed.ReadFrom(gzReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Decompressed Data")	
	fmt.Printf("%v\n", decompressed.Bytes())	

	
}