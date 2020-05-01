package main

import (
	"fmt"
)

var y = 100
var z string = "shaken, not stirred"

// vars can be declared outside the scope of a function

func main() {
	x := 42
	// declares a variable and assigns a value to it
	fmt.Println(x)
	x = 99
	// after a variable is already declared, you can only assign new values to it
	fmt.Println(x)

	fmt.Println(y)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
}
