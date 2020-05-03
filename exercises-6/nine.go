package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4}

	fmt.Println(product(x...))

	fmt.Println(evens(product, x...))
}

func product(listOfNumbers ...int) int {
	total := 1

	for _, v := range listOfNumbers {
		total *= v
	}

	return total
}

func evens(fu func(listOfNumbers ...int) int, listOfNumbers ...int) int {
	var evenNumbers []int

	for _, v := range listOfNumbers {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}
	}

	return fu(evenNumbers...)
}
