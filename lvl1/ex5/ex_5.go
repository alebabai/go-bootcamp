package ex5

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

// Resolve is a function that provide solution of exercise 5 ninja-level-1
func Resolve() {
	fmt.Println("Exercise 5")

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

	fmt.Println()
}
