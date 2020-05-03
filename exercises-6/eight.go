package main

import (
	"fmt"
)

func main() {
	x := returnsFunc()

	fmt.Println(x())
}

func returnsFunc() func() int {
	return func() int {
		return 1
	}
}
