package main

import (
	"fmt"
	"log"
	"net/http"
)

type Animal interface {
	IsPet() bool
	GetInfo() string
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
	http.HandleFunc("/", getFamilyInfo)      // set router
	err := http.ListenAndServe(":8000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getFamilyInfo(w http.ResponseWriter, r *http.Request) {
	j := Human{"Jonathan", "Thom", 28}
	l := Human{"Laura", "Syvertson", 28}
	e := Dog{"Ernie", "The Dog", 8}

	family := Family{Members: []Animal{j, l, e}}

	family.totalUpPetCount()
	fmt.Fprintf(w, "Family pet count: %v\n\n", family.PetCount["Pet Count"])
	fmt.Fprintf(w, "Family non-pet count: %v\n\n", family.PetCount["Non Pet Count"])
	family.PrintPeopleInfo(w)
}

func (d Dog) IsPet() bool {
	return true
}

func (d Human) IsPet() bool {
	return false
}

func (family *Family) totalUpPetCount() {
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

func (family Family) PrintPeopleInfo(w http.ResponseWriter) {
	for _, m := range family.Members {
		info := m.GetInfo()
		fmt.Fprintf(w, info)
	}
}

func (h Human) GetInfo() string {
	return fmt.Sprintf("%s %s is %v years old\n", h.FirstName, h.LastName, h.Age)
}

func (d Dog) GetInfo() string {
	dogYears := d.Age * 7
	return fmt.Sprintf("%s %s is %v dog years old\n", d.FirstName, d.LastName, dogYears)
}

// go build
// ./build-a-gopher
// localhost:8000
// Family pet count: 1

// Family non-pet count: 2

// Jonathan Thom is 28 years old
// Laura Syvertson is 28 years old
// Ernie The Dog is 56 dog years old
