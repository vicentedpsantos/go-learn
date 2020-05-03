package main

import (
	"fmt"
)

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	vicente = person{
		firstName:        "Vicente",
		lastName:         "Santos",
		favoriteIceCream: string{"Chocolate", "Mint", "Lemon"},
	}

	tamires = person{
		firstName:        "Tamires",
		lastName:         "Quito",
		favoriteIceCream: string{"Chocolate", "Grapes", "Ninho"},
	}

	personsMap := map[string]person{
		vicente.lastName:  vicente,
		tamirers.lastName: tamires,
	}
}
