package main

import (
	"fmt"
)

func main() {
	xi := []int{10, 23, 42, 64, 32, 12, 65, 21, 87, 74, 92, 82, 12, 23}

	totalSum := sum(xi...)
	fmt.Println(totalSum)

	totalSumEvens := sumsEven(sum, xi...)
	fmt.Println(totalSumEvens)

	totalSumOdds := sumsOdds(sum, xi...)
	fmt.Println(totalSumOdds)
}

func sum(listOfNumbers ...int) int {
	total := 0

	for _, v := range listOfNumbers {
		total += v
	}

	return total
}

func sumsEven(f func(listOfNumbers ...int) int, listOfNumbers ...int) int {
	var evenNumbers []int

	for _, v := range listOfNumbers {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}
	}

	return f(evenNumbers...)
}

func sumsOdds(f func(listOfNumbers ...int) int, listOfNumbers ...int) int {
	var oddNumbers []int

	for _, v := range listOfNumbers {
		if v%2 != 0 {
			oddNumbers = append(oddNumbers, v)
		}
	}

	return f(oddNumbers...)
}
