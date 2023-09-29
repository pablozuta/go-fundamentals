// sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier is an interface that defines notification type behavior
type notifier interface {
	notify()
}

// user defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending user email to %s <%s>\n", u.name, u.email)
}

// admin defines an admin in the program
type admin struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver
func (a *admin) notify() {
	fmt.Printf("Sending user email to %s <%s>\n", a.name, a.email)
}

func main() {
	// create an user value and pass it to sendNotification
	bill := user{"Bill", "bill.oceanavenue@gmail.com"}
	sendNotification(&bill)

	// create an admin value and pass it to sendNotification
	tom := user{"Tom", "tom.red1652.com"}
	sendNotification(&tom)
}

// sendNotification accepts values that implement the notifier interface and sends notifications
func sendNotification(n notifier) {
	n.notify()
}
