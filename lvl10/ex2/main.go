package ex2

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 2 ninja-level-10
func Resolve() {
	fmt.Println("Exercise 2")
	step1()
	step2()
	fmt.Println()
}

func step1() {
	fmt.Println("Step 1")
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func step2() {
	fmt.Println("Step 2")
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
