package ex8

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 8 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 8")

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for key, slice := range m {
		fmt.Println("This is the record for", key)
		for index, value := range slice {
			fmt.Println("\t", index, value)
		}
	}

	fmt.Println()
}
