package ex5

import (
	"fmt"
	"math"
)

type rectangle struct {
	l, w float64
}

func (r rectangle) area() float64 {
	return r.l * r.w
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2.0)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("Area of %T is %v\n", s, s.area())
}

// Resolve is a function that provide solution of exercise 5 ninja-level-6
func Resolve() {
	fmt.Println("Exercise 5")

	r := rectangle{l: 13, w: 20}
	c := circle{r: 13}

	shapes := []shape{r, c}
	for _, s := range shapes {
		info(s)
	}

	fmt.Println()
}
