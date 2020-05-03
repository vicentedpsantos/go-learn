package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Println("Speaking from a person")
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("Speaking from a secret agent")
}

func (sa secretAgent) cry() {
	fmt.Println("=(")
}

type human interface {
	speak()
	cry()
}

func isHuman(a interface{}) bool {
	switch a.(type) {
	case human:
		return true
	default:
		return false
	}
}

func main() {
	p1 := person{
		firstName: "Vicente",
		lastName:  "Santos",
		age:       29,
	}

	jb := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       76,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(jb)

	fmt.Println(isHuman(p1))
	fmt.Println(isHuman(jb))
}
