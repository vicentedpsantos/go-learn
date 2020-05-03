package main

import (
	"fmt"
)

func main() {
	aInt := foo()
	fmt.Println(aInt)

	anotherInt, aString := bar()
	fmt.Println(anotherInt, aString)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "aString"
}
