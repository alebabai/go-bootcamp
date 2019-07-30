package ex4

import (
	"fmt"
)

type nagual int

var x nagual

// Resolve is a function that provide solution of exercise 4 ninja-level-1
func Resolve() {
	fmt.Println("Exercise 4")

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	fmt.Println()
}
