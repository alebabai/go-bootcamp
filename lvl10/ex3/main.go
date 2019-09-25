package ex3

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 3 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 3")
	c := gen()
	receive(c)
	fmt.Println()
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("Value from channel", v)
	}
}
