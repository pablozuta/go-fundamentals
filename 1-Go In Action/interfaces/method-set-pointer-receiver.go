// sample program to show how to use an interface in go

package main

import "fmt"

// notifier is an interface that defined notification type behavior
type notifier interface {
	notify()
}

// user defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending user email to: %s <%s>\n", u.name, u.email)

}

func main() {
	// create a value of type user and send a notification
	u := user{"neo", "neo.biology.matrix@gmail.com"}
	// usamos signo de pointer para no obtener error de compilacion
	sendNotification(&u)

}

// sendNotification  accepts values that implement the notifier interface and send notifications.
func sendNotification(n notifier) {
	n.notify()
}
