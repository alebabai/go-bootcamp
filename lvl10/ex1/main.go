package ex1

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 1 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 1")

	fmt.Println("Using func literal, aka, anonymous self-executing func")
	c1 := make(chan int)
	go func() {
		c1 <- 42
	}()
	fmt.Println(<-c1)

	fmt.Println("Using a buffered channel")
	c2 := make(chan int, 1)
	c2 <- 42
	fmt.Println(<-c2)

	fmt.Println()
}
