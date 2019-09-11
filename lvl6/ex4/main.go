package ex4

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

// Resolve is a function that provide solution of exercise 4 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 4")
	
	p := person{
		first: "Ivan",
		last:  "Ivanov",
		age:   34,
	}
	p.speak(

	)
	fmt.Println()
}
