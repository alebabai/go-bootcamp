package ex5

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 5 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 5")

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Origin", slice)

	slice = append(slice[:3], slice[6:]...)
	fmt.Println("After delete&slicing", slice)

	fmt.Println()
}
