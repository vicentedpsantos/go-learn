package main

import (
	"fmt"
)

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}

type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	s1 := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: true,
	}

	t1 := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "Black",
		},
		fourWheel: true,
	}

	fmt.Println(s1)
	fmt.Println(t1)

	fmt.Println(s1.doors) // Promoted values
	fmt.Println(t1.fourWheel)
}
