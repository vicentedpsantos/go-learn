package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := int64(42)

	fmt.Println(x)
	fmt.Println(strconv.FormatInt(x, 2))

	hex := fmt.Sprintf("%X", x)
	fmt.Println(hex)

	fmt.Printf("%d \t%b \t%x", x, x, x)
}
