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

	fmt.Println(people)
}

// go build
// ./build-a-gopher
// [{Jonathan Thom 28} {Laura Syvertson 28} {Ernie The Dog 8}]
