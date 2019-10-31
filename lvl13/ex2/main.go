package ex2

import (
	"fmt"

	"github.com/alebabai/go-bootcamp/lvl13/ex2/quote"
	"github.com/alebabai/go-bootcamp/lvl13/ex2/word"
)

type canine struct {
	name string
	age  int
}

// Resolve is a function that provide solution of exercise 2 ninja-level-13
func Resolve() {
	fmt.Println("Exercise 2")

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	fmt.Println()
}
