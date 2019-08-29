package ex2

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 2 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 2")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Printf("Index: %d\tvalue: %v\n", i, v)
	}

	fmt.Printf("Slice's type: %T", slice)

	fmt.Println()
}
