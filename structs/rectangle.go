package main
import "fmt"

type Rectangle struct {
	width float64
	height float64
}

func(r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	rectangle := Rectangle{width: 10, height: 20}
	fmt.Println("Area: ", rectangle.area())
}