package main

import (
	"fmt"
)

func main() {
	fmt.Println("Helloe everyone, this is a go program")
	foo()
	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am in foo")
}
