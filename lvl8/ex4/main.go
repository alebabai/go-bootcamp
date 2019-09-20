package ex4

import (
	"fmt"
	"sort"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-8
func Resolve() {
	fmt.Println("Exercise 4")

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println()
}
