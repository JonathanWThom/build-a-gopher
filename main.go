package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p := Person{"Jonathan", "Thom", 28}
	fmt.Printf("The person %s %s is %v", p.FirstName, p.LastName, p.Age)
}

// go build
// ./jonathan
// The person Jonathan Thom is 28
