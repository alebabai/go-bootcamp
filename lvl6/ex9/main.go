package ex9

import (
	"fmt"
)

func foo(action func(...interface{}), values ...interface{}) {
	for _, v := range values {
		action(v)
	}
}

// Resolve is a function that provide solution of exercise 9 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 9")

	printer := func(value ...interface{}) {
		fmt.Println(value...)
	}

	foo(printer, 1, 2, 3, "Fight", "Club")

	fmt.Println()
}
