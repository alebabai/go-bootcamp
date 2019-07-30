package ex3

import (
	"fmt"
)

var x int
var y string
var z bool

// Resolve is a function that provide solution of exercise 3 ninja-level-1
func Resolve() {
	fmt.Println("Exercise 3")

	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
	fmt.Println()
}
