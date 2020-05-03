package main

import "fmt"

func main() {
	a := make([]int, 10)

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", a)
}
