package ex4

import (
	"fmt"
	"time"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 4")

	const startYear = 1995
	currentYear := time.Now().Year()
	year := startYear

	for {
		if year > currentYear {
			break
		}
		fmt.Printf("I lived in the %d\n", year)
		year++
	}

	fmt.Println()
}
