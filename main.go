package main

import "fmt"

func main() {
	firstName := "Jonathan"
	fmt.Printf("Hello %s\n", firstName)

	var lastName string
	lastName = "Thom"
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}

// go build
// ./jonathan
// Hello Jonathan
// Hello Jonathan Thom
