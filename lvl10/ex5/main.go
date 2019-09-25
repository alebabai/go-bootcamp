package ex5

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 5 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 5")

	c := make(chan int)

	go func() {
		c <- 1
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

	fmt.Println()
}
