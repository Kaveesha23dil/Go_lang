package main

import "fmt"

// -----------------------------------
// Go Variables: var vs :=
// -----------------------------------

func main() {

	// -----------------------------------
	// 1. Using var
	// -----------------------------------

	// Declaring a variable with data type
	var name string

	// Assigning a value later
	name = "Alice"

	// Declaring and initializing together
	var rating float64 = 4.5

	// Type inference with var
	// Go automatically detects the type
	var city = "Colombo"
	var price = 150

	// -----------------------------------
	// 2. Using :=
	// -----------------------------------

	// Short variable declaration
	// Go automatically:
	// - creates the variable
	// - detects the data type

	age := 23
	isStudent := true

	// -----------------------------------
	// Printing Values
	// -----------------------------------

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Rating:", rating)
	fmt.Println("City:", city)
	fmt.Println("Price:", price)
	fmt.Println("Is Student:", isStudent)
}

/*

===================================
IMPORTANT NOTES
===================================

1. var Syntax
----------------
var variableName dataType = value

Example:
var name string = "Alice"

2. var with Type Inference
----------------
var city = "Colombo"

Go automatically detects:
city -> string

3. := Syntax
----------------
variableName := value

Example:
age := 23

4. := Rules
----------------
- Only works inside functions
- Cannot be used for global variables

5. Best Practice
----------------
Go developers usually prefer := inside functions
because it is shorter and cleaner.

*/
