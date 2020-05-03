package main

import (
	"fmt"
)

func main() {
	x := func(s string) string {
		return fmt.Sprintf("Using an %s function", s)
	}("Anonymous")

	fmt.Println(x)
}
