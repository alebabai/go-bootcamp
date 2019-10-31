package ex3

import (
	"fmt"

	"github.com/alebabai/go-bootcamp/lvl13/ex3/math"
)

type canine struct {
	name string
	age  int
}

// Resolve is a function that provide solution of exercise 3 ninja-level-13
func Resolve() {
	fmt.Println("Exercise 3")

	xxi := gen()
	for _, v := range xxi {
		fmt.Println(math.CenteredAvg(v))
	}

	fmt.Println()
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
