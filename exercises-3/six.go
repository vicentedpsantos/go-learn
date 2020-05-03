package main

import (
	"fmt"
)

func main() {
	x := 42

	if x < 43 {
		fmt.Println("This will print")
	} else if x > 42 {
		fmt.Println("This will not print")
	} else {
		fmt.Println("Wont print")
	}
}
