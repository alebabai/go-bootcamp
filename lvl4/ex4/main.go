package ex4

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 4")

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Origin", slice)

	slice = append(slice, 52)
	fmt.Println("Append single value", slice)

	slice = append(slice, 53, 54, 55)
	fmt.Println("Append multiple values", slice)

	anotherSlice := []int{56, 57, 58, 59, 60}
	slice = append(slice, anotherSlice...)
	fmt.Println("Append another slice", slice)

	fmt.Println()
}
