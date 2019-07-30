package ex6

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 6 ninja-level-2
func Resolve() {
	fmt.Println("Exercise 6")

	const (
		currentYear = 2019
		a           = currentYear + iota
		b           = currentYear + iota
		c           = currentYear + iota
		d           = currentYear + iota
	)

	fmt.Println(a, b, c, d)

	fmt.Println()
}
