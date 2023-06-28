package main

import "fmt"

func main() {
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)

	day := Friday

	switch day {
	case Sunday:
		fmt.Println("Today is Sunday, day", Sunday, "of the week.")
	case Monday:
		fmt.Println("Today is Monday, day", Monday, "of the week.")
	case Tuesday:
		fmt.Println("Today is Tuesday, day", Tuesday, "of the week.")
	case Wednesday:
		fmt.Println("Today is Wednesday, day", Wednesday, "of the week.")
	case Thursday:
		fmt.Println("Today is Thursday, day", Thursday, "of the week.")
	case Friday:
		fmt.Println("Today is Friday, day", Friday, "of the week.")
	case Saturday:
		fmt.Println("Today is Saturday, day", Saturday, "of the week.")
	}
}
