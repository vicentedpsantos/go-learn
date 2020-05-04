package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs := []byte(s)

	p, err := bcrypt.GenerateFromPassword(bs, bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)

	loginPassword := `password1233`

	err = bcrypt.CompareHashAndPassword(p, []byte(loginPassword))

	if err != nil {
		fmt.Println("Password is incorrect")
	} else {
		fmt.Println("Password is correct")
	}
}
