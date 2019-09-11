package ex2

import (
	"fmt"
)

func foo(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func bar(numbers []int) int {
	return foo(numbers...)
}

// Resolve is a function that provide solution of exercise 2 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 2")

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(numbers...), bar(numbers))

	fmt.Println()
}
