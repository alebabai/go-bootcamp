package ex7

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 7 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 7")

	f := func() {
		fmt.Println("Hello from anonymous func")
	}

	f()

	fmt.Println()
}
