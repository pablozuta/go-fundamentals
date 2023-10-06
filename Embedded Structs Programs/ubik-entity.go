// In this program, there are four struct types: entity, object, ubikSubject, and ubikEvent. The object struct embeds the entity struct using the syntax entity. The ubikSubject struct also embeds the entity struct using the syntax entity. The ubikEvent struct embeds both the object and ubikSubject structs using the syntax object and ubikSubject.

// In the `main` function, a new object instance is created with a name of "Tin Can" and a location of "Moon Base Alpha". A new ubikSubject instance is also created with a name of "Pat Conley" and a condition of "half-life". A new ubikEvent instance is then created with the object and ubikSubject instances as its embedded structs and a time of "3:00 PM". The program then prints out the ubikEvent information using `fmt.Printf`.
package main

import "fmt"

type entity struct {
	name string
}

type object struct {
	entity
	location string
}

type ubikSubject struct {
	entity
	condition string
}
type ubikEvent struct {
	object
	ubikSubject
	time string
}

func main() {
	// create a new object instance
	tinCan := object{
		entity:   entity{name: "Tin Can"},
		location: "Moon Base Alpha.",
	}

	// create a ubikSubject instance
	pat := ubikSubject{
		entity:    entity{name: "Pat Conley"},
		condition: "half-life",
	}

	// create a new ubikEvent instance
	ubikEvent := ubikEvent{
		object:      tinCan,
		ubikSubject: pat,
		time:        "3:00 PM",
	}

	// Print out the UbikEvent information
	fmt.Printf("Event Name: %s\n", ubikEvent.object.name)
	fmt.Printf("Object Location: %s\n", ubikEvent.location)
	fmt.Printf("Subject Condition: %s\n", ubikEvent.condition)
	fmt.Printf("Event Time: %s\n", ubikEvent.time)
}
