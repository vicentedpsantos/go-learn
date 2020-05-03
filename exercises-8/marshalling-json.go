package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "Vicente",
		Age:  29,
	}

	p2 := Person{
		Name: "Tamires",
		Age:  29,
	}

	people := []Person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
