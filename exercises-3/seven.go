package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Wont print")
	case true:
		fmt.Println("Will print")
		fallthrough
	case true:
		fmt.Println("Will print")
	default:
		fmt.Println("Wont print")
	}

	caseWithValue("Gremio")
}

func caseWithValue(value string) {
	switch value {
	case "Bond":
		fmt.Println("James Bond")
	case "Gremio":
		fmt.Println("Multichampion")
	}
}
