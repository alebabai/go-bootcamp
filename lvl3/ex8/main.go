package ex8

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 8 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 8")

	switch {
	case true:
		fmt.Println("Yeah!")
	case false:
		fmt.Println("Will never be invoked")
	}

	fmt.Println()
}
