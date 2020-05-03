package main

import (
	"fmt"
)

func main() {
	s := foo()

	fmt.Println(s)

	s1 := buzz()

	fmt.Println(s1())
	fmt.Println(s1) // return function
}

func foo() string {
	return "Bar"
}

func buzz() func() string {
	return func() string {
		return "Fizz"
	}
}
