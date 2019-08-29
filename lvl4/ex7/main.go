package ex7

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 7 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 7")

	slice := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(slice)

	for i, row := range slice {
		fmt.Println("Raw index: ", i)
		for j, value := range row {
			fmt.Println("Column index: ", j)
			fmt.Println("Value: ", value)
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println()
}
