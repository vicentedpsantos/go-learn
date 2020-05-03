package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string][]string)

	myMap["bond_james"] = []string{"shaken, not stirred", "martini", "women"}
	myMap["moneypenny_miss"] = []string{"James Bond", "literature", "computer science"}
	myMap["no_dr"] = []string{"being evil", "ice cream", "sunsets"}

	for k, v := range myMap {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i+1, v2)
		}
	}
}
