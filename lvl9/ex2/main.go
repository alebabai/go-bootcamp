package ex2

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("My name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

// Resolve is a function that provide solution of exercise 2 ninja-level-9
func Resolve() {
	fmt.Println("Exercise 2")

	p := person{"Tyler Durden"}
	saySomething(&p)

	// saySomething(p) //doesn't work

	p.speak()
	(&p).speak()

	fmt.Println()
}
