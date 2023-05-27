package main

import "fmt"

type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

func (c *Counter) Value() int {
    return c.count
}

func main() {
    // Create a new Counter object
    counter := Counter{10}

    // Increment the counter en uno
    counter.Increment()

    // Get the current value of the counter
    value := counter.Value()

    // Print the value (11)
    fmt.Println("Counter value:", value)
}

