package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return fmt.Sprintf("Hi, my name is %s and my age is %d!", p.first, p.age)
}

func main() {
	p := person{
		first: "Vicente",
		last:  "Santos",
		age:   29,
	}

	fmt.Println(p.speak())
}
