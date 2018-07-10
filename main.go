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

	family.totalUpPetCount()
	fmt.Printf("Family pet count: %v\n\n", family.PetCount["Pet Count"])
	fmt.Printf("Family non-pet count: %v\n\n", family.PetCount["Non Pet Count"])

	family.PrintPeopleInfo()
}

func (family Family) totalUpPetCount() {
	pets := make(map[string]int)

	for _, p := range family.People {
		if p.IsPet == true {
			pets["Pet Count"]++
		} else if p.IsPet == false {
			pets["Non Pet Count"]++
		}
	}

	family.PetCount = pets
}

func (family Family) PrintPeopleInfo() {
	for _, p := range family.People {
		p.PrintInfo()
	}
}

func (p Person) PrintInfo() {
	fmt.Printf("%s %s is %v years old\n", p.FirstName, p.LastName, p.Age)
}

// go build
// ./build-a-gopher
// Family pet count: 0

// Family non-pet count: 0

// Jonathan Thom is 28 years old
// Laura Syvertson is 28 years old
// Ernie The Dog is 8 years old
