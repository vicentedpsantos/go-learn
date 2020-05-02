package main

import (
	"fmt"
)

const (
	a = iota + 2017
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
