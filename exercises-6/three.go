package main

import (
	"fmt"
)

func main() {
	defer deferedFunc()
	aFunc()
}

func aFunc() {
	fmt.Println("Running before...")
}

func deferedFunc() {
	fmt.Println("Running defered...")
}
