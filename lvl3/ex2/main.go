package ex2

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 2 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 2")

	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	fmt.Println()
}
