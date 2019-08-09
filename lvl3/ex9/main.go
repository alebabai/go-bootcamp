package ex9

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 9 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 9")

	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}

	fmt.Println()
}
