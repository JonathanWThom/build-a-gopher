package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	j := Person{"Jonathan", "Thom", 28}
	l := Person{"Laura", "Syvertson", 28}

	people := []Person{j, l}
	e := Person{"Ernie", "The Dog", 8}
	people = append(people, e)

	for i, p := range people {
		fmt.Printf("%s %s is %v years old\n", p.FirstName, p.LastName, p.Age)
		fmt.Printf("They are number %v is this list\n\n", i)
	}
}

// go build
// ./build-a-gopher

// Jonathan Thom is 28 years old
// They are number 0 is this list
//
// Laura Syvertson is 28 years old
// They are number 1 is this list
//
// Ernie The Dog is 8 years old
// They are number 2 is this list
