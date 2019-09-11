package ex6

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 6 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 6")

	func() {
		fmt.Println("Hello from anonymous func")
	}()

	fmt.Println()
}
