package main

import (
	"fmt"
)

func main() {
	var city string
	city = "Colombo"

	var channel string = "Tech With Tim" //inferred to string

	// := is a short variable declaration, it can only be used inside functions and it infers the type of the variable
	country := "Sri Lanka"

	fmt.Println("City:", city)
	fmt.Println("Channel:", channel)
	fmt.Println("Country:", country)
}
