package main

import (
	"fmt"
)

func main() {
	x := struct {
		firstName string
		lastName  string
		age       int
		favDrinks []string
	}{
		firstName: "Vicente",
		lastName:  "Santos",
		age:       29,
		favDrinks: []string{"Coke", "Beer", "Orange Juice"},
	}

	fmt.Println(x)
}
