package ex4

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-2
func Resolve() {
	fmt.Println("Exercise 4")

	a := 13
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)

	fmt.Println()
}
