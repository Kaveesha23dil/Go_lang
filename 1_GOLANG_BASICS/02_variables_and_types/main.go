package main

import "fmt"

func main() {

	// Declaring a variable of type string
	var name string

	// Assigning a value to the variable
	name = "Alice"

	// Declaring and initializing a variable in one line
	age := 23

	// Printing the values of the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	var ratting float64 = 4.5
	fmt.Println("Ratting:", ratting)
}

// 1. Using var
// var is the traditional way to declare variables.

// Syntax -> var variableName dataType = value

//Type Inference with var
//Go can automatically detect the type.

//var city = "Colombo"
//var price = 150

// 2. Using :=
// This is the short variable declaration syntax.

// Syntax ->  variableName := value

// Example:
//        name := "Alice"
//        age := 23

//Go automatically:
//          creates variable
//          detects type

//Best Practice in Go
//Go developers usually prefer:
//                             :=
//inside functions because it is cleaner.
