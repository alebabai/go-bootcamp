package ex2

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Tyler Durden"
}

// Resolve is a function that provide solution of exercise 2 ninja-level-7
func Resolve() {
	fmt.Println("Exercise 2")

	p1 := person{
		name: "The storyteller",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

	fmt.Println()
}
