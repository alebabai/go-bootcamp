package ex6

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 6 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 6")

	c := make(chan int)

	go func(cs chan<- int) {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(cs)
	}(c)

	for v := range c {
		fmt.Println("Value from channel", v)
	}

	fmt.Println()
}
