package main

import (
	"fmt"
)

type person struct {
	fullName string
	age      int
}

func (p *person) speak() {
	fmt.Println("Person speaks")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	vicente := person{
		fullName: "Vicente Santos",
		age:      29,
	}

	saySomething(&vicente) //works

	saySomething(vicente) // breaks -> cannot pass type person into a func that receives a pointer to person
}
