package main

import (
	"fmt"
)

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	fmt.Println(x)

	for i, v := range x {
		fmt.Printf("\nThis is the array number %d:\n\t %v", i+1, v)

		for j, z := range v {
			fmt.Printf("\n\t\tThis is the value number %d of the array number %d: %v", j+1, i+1, z)
		}
	}
}
