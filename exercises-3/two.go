package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 26; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U \n", toChar(i))
		}
	}
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}
