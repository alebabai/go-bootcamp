package ex3

import (
	"fmt"
)

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

// Resolve is a function that provide solution of exercise 3 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 3")

	defer foo()
	fmt.Println("after DEFER foo")

	fmt.Println()
}
