package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Full Name:", fullName)

	// String functions
	fmt.Println("Length of Full Name:", len(fullName))
	fmt.Println("Uppercase Full Name:", strings.ToUpper(fullName))
	fmt.Println("Lowercase Full Name:", strings.ToLower(fullName))
	fmt.Println("Contains 'John':", strings.Contains(fullName, "John"))
	fmt.Println("Index of 'Doe':", strings.Index(fullName, "Doe"))
}
