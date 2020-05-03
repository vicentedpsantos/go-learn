package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string][]string)

	myMap["bond_james"] = []string{"shaken, not stirred", "martini", "women"}
	myMap["moneypenny_miss"] = []string{"James Bond", "literature", "computer science"}
	myMap["no_dr"] = []string{"being evil", "ice cream", "sunsets"}

	presentMap(myMap)

	myMap["santos_vicente"] = []string{"gremio", "tamires", "family"}

	presentMap(myMap)

	delete(myMap, "no_dr")

	presentMap(myMap)
}

func presentMap(aMap map[string][]string) {
	for k, v := range aMap {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i+1, v2)
		}
	}
}
