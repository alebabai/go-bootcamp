package ex1

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 1 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 1")

	arr := [5]int{1, 2, 3, 4, 13}

	for i, v := range arr {
		fmt.Printf("Index: %d\tvalue: %v\n", i, v)
	}

	fmt.Printf("Array's type: %T", arr)

	fmt.Println()
}
