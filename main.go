package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	IsPet     bool
}

type Family struct {
	People   []Person
	PetCount map[string]int
}

func main() {
	j := Person{"Jonathan", "Thom", 28, false}
	l := Person{"Laura", "Syvertson", 28, false}
	e := Person{"Ernie", "The Dog", 8, true}

	family := Family{People: []Person{j, l, e}}

	totalUpPetCount(&family)
	fmt.Printf("Family pet count: %v\n\n", family.PetCount["Pet Count"])
	fmt.Printf("Family non-pet count: %v\n\n", family.PetCount["Non Pet Count"])
}

func totalUpPetCount(family *Family) {
	pets := make(map[string]int)

	for _, p := range family.People {
		fmt.Printf("%s %s is %v years old\n", p.FirstName, p.LastName, p.Age)

		if p.IsPet == true {
			pets["Pet Count"]++
		} else if p.IsPet == false {
			pets["Non Pet Count"]++
		}
	}

	family.PetCount = pets
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
