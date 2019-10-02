package ex3

import (
	"fmt"
	"log"
)

type customError struct {
	cause string
}

func (e customError) Error() string {
	return fmt.Sprintf("Error happend due to %v", e.cause)
}

func handleError(e error) {
	log.Printf("Handle error: %v ", e)
}

// Resolve is a function that provide solution of exercise 3 ninja-level-11
func Resolve() {
	fmt.Println("Exercise 3")

	e := customError{"some magic thing"}
	handleError(e)

	fmt.Println()
}
