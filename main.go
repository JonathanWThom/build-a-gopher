package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	IsPet     bool
}

func main() {
	j := Person{"Jonathan", "Thom", 28, false}
	l := Person{"Laura", "Syvertson", 28, false}

	people := []Person{j, l}
	e := Person{"Ernie", "The Dog", 8, true}
	people = append(people, e)

	pets := make(map[string]int)

	for _, p := range people {
		fmt.Printf("%s %s is %v years old\n", p.FirstName, p.LastName, p.Age)

		if p.IsPet == true {
			pets["Pet Count"]++
		} else if p.IsPet == false {
			pets["Non Pet Count"]++
		}
	}

	fmt.Printf("Family pet count: %v\n\n", pets["Pet Count"])
	fmt.Printf("Family non-pet count: %v\n\n", pets["Non Pet Count"])
}

// go build
// ./build-a-gopher

// Jonathan Thom is 28 years old
//
// Laura Syvertson is 28 years old
//
// Ernie The Dog is 8 years old
// Family pet count: 1
//
// Family non-pet count: 2
