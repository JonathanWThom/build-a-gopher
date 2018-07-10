package main

import "fmt"

type Animal interface {
	IsPet() bool
	PrintInfo()
}

type Dog struct {
	FirstName string
	LastName  string
	Age       int
}

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

type Family struct {
	Members  []Animal
	PetCount map[string]int
}

func main() {
	j := Human{"Jonathan", "Thom", 28}
	l := Human{"Laura", "Syvertson", 28}
	e := Dog{"Ernie", "The Dog", 8}

	family := Family{Members: []Animal{j, l, e}}

	family.totalUpPetCount()
	fmt.Printf("Family pet count: %v\n\n", family.PetCount["Pet Count"])
	fmt.Printf("Family non-pet count: %v\n\n", family.PetCount["Non Pet Count"])

	family.PrintPeopleInfo()
}

func (d Dog) IsPet() bool {
	return true
}

func (d Human) IsPet() bool {
	return false
}

func (family Family) totalUpPetCount() {
	pets := make(map[string]int)

	for _, m := range family.Members {
		if m.IsPet() == true {
			pets["Pet Count"]++
		} else if m.IsPet() == false {
			pets["Non Pet Count"]++
		}
	}

	family.PetCount = pets
}

func (family Family) PrintPeopleInfo() {
	for _, m := range family.Members {
		m.PrintInfo()
	}
}

func (h Human) PrintInfo() {
	fmt.Printf("%s %s is %v years old\n", h.FirstName, h.LastName, h.Age)
}

func (d Dog) PrintInfo() {
	dogYears := d.Age * 7
	fmt.Printf("%s %s is %v dog years old\n", d.FirstName, d.LastName, dogYears)
}

// go build
// ./build-a-gopher
// Family pet count: 0

// Family non-pet count: 0

// Jonathan Thom is 28 years old
// Laura Syvertson is 28 years old
// Ernie The Dog is 56 dog years old
