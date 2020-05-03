package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4}

	fmt.Println(foo(xi...))

	fmt.Println(bar(xi))
}

func foo(listOfNumbers ...int) int {
	return sum(listOfNumbers...)
}

func bar(xi []int) int {
	return sum(xi...)
}

func sum(listOfNumbers ...int) int {
	total := 0

	for _, v := range listOfNumbers {
		total += v
	}

	return total
}
