package ex1

import (
	"fmt"
)

func foo() int {
	return 13
}

func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}

// Resolve is a function that provide solution of exercise 1 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 1")

	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)

	fmt.Println()
}
