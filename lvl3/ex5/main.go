package ex5

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 5 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 5")

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

	fmt.Println()
}
