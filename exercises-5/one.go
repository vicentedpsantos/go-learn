package main

import (
	"fmt"
)

type person struct {
	firstName        string
	lastName         string
	age              int
	favoriteIceCream []string
}

func main() {
	vicente := person{
		firstName:        "Vicente",
		lastName:         "Santos",
		age:              29,
		favoriteIceCream: []string{"Chocolate", "Mint"},
	}

	tamires := person{
		firstName:        "Tamires",
		lastName:         "Quito",
		age:              29,
		favoriteIceCream: []string{"Chocolate", "Grape"},
	}

	presentIceCreams(vicente)
	presentIceCreams(tamires)
}

func presentIceCreams(person person) {
	fmt.Printf("%s's favorite ice creams are: \n", person.firstName)

	for k, v := range person.favoriteIceCream {
		fmt.Printf("%d \t - %s\n", k+1, v)
	}
}
