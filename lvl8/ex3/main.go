package ex3

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Resolve is a function that provide solution of exercise 3 ninja-level-8
func Resolve() {
	fmt.Println("Exercise 3")

	users := []user{
		user{
			Name: "Tyler Durden",
			Age:  32,
		},
		user{
			Name: "The storyteller",
			Age:  34,
		},
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(users)

	fmt.Println()
}
