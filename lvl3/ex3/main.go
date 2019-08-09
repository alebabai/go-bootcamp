package ex3

import (
	"fmt"
	"time"
)

// Resolve is a function that provide solution of exercise 3 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 3")

	const startYear = 1995
	currentYear := time.Now().Year()
	year := startYear

	for year <= currentYear {
		fmt.Printf("I lived in the %d\n", year)
		year++
	}

	fmt.Println()
}
