package ex7

import (
	"fmt"
	"math/rand"
)

// Resolve is a function that provide solution of exercise 7 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 7")

	const gsCount = 10
	const numbersCount = 10

	c := make(chan int)

	for i := 0; i < gsCount; i++ {
		go func(cs chan<- int) {
			for j := 0; j < numbersCount; j++ {
				c <- rand.Int()
			}
		}(c)
	}

	for i := 0; i < gsCount*numbersCount; i++ {
		v, ok := <-c
		fmt.Println("Value from channel :", v, "\t", "is it okay :", ok)
	}

	fmt.Println()
}
