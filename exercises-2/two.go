package main

import (
	"fmt"
)

func main() {
	g := 1 == 2
	h := 32 <= 211
	i := 32 >= 21
	j := "String a" != "String a"
	k := 3 < 4
	l := 4 > 3

	fmt.Println(g, h, i, j, k, l)
}
