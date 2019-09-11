package ex8

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 8 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 8")

	f := func() func() {
		return func() {
			fmt.Println("Hello from anonymous func")
		}
	}()

	f()

	fmt.Println()
}
