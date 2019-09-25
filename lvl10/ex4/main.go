package ex4

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 4")
	q := make(chan int)
	c := gen(q)
	receive(c, q)
	fmt.Println()
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}
