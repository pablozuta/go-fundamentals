package main  
  
import (  
    "fmt"  
    "unsafe"  
)  
  
type Person struct {  
    name string  
    age int  
}  
  
func main() {  
    // Create a Person struct and convert it to a byte slice  
    p := Person{name: "John", age: 30}  
    b := *(*[]byte)(unsafe.Pointer(&p))  

    // Convert the byte slice back to a Person struct  
    var p2 Person  
    p2 = *(*Person)(unsafe.Pointer(&b[0]))  
 
    // Print out the name and age of the Person struct  
    fmt.Println("Name:", p2.name)  
    fmt.Println("Age:", p2.age)  
} 