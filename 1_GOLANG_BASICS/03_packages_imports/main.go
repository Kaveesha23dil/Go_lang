package main

// imports brings external packages into our code, allowing us to use their functionality.
import (
	"fmt"
	"math"
)

func main() {
	// packegeName.FunctionName() - syntax to call a function from an imported package

	fmt.Println("sqrt(25)", math.Sqrt(25))
}
