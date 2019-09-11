package ex10

import (
	"fmt"
)

func makeCounter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}

// Resolve is a function that provide solution of exercise 10 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 10")

	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println()
}
