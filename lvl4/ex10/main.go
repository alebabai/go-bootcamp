package ex10

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 10 ninja-level-4
func Resolve() {
	fmt.Println("Exercise 10")

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	delete(m, "no_dr")

	for key, slice := range m {
		fmt.Println("This is the record for", key)
		for index, value := range slice {
			fmt.Println("\t", index, value)
		}
	}

	fmt.Println()
}
