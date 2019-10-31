package ex1

import (
	"fmt"

	"github.com/alebabai/go-bootcamp/lvl13/ex1/dog"
)

type canine struct {
	name string
	age  int
}

// Resolve is a function that provide solution of exercise 1 ninja-level-13
func Resolve() {
	fmt.Println("Exercise 1")

	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))

	fmt.Println()
}
