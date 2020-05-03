package main

import (
	"fmt"
)

type person struct {
	firstName string
}

func changeMe(p *person, newfirstName string) {
	(*p).firstName = newfirstName
}

func main() {
	vicente := person{
		firstName: "Vicente",
	}

	fmt.Println(vicente.firstName)

	changeMe(&vicente, "Daniel")

	fmt.Println(vicente.firstName)
}
