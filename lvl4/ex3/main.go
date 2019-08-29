package ex3

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 3 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 3")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)

	fmt.Println()
}
