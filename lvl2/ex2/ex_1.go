package ex2

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 2 ninja-level-2
func Resolve() {
	fmt.Println("Exercise 2")

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)

	fmt.Println()
}
