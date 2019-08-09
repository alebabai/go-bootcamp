package ex7

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 7 ninja-level-3
func Resolve() {
	fmt.Println("Exercise 7")

	lang := "java"

	if lang == "go" {
		fmt.Println("Your will get well paid job")
	} else if lang == "java" {
		fmt.Println("Your will write a lot of Plain Old Java Objects")
	} else {
		fmt.Println("Angular developer")
	}

	fmt.Println()
}
