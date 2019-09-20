package ex1

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Resolve is a function that provide solution of exercise 1 ninja-level-8
func Resolve() {
	fmt.Println("Exercise 1")

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

	json, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))

	fmt.Println()
}
