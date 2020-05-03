package main

import (
	"fmt"
)

func main() {
	bornIn := 1990

	for bornIn < 2020 {
		fmt.Printf("I was already born in %d \n", bornIn)
		bornIn++
	}
}
