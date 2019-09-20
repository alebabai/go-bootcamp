package ex2

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Resolve is a function that provide solution of exercise 2 ninja-level-8
func Resolve() {
	fmt.Println("Exercise 2")

	rawJSON := `[{"name":"Tyler Durden","age":32},{"name":"The storyteller","age":34}]`

	var users []user

	err := json.Unmarshal([]byte(rawJSON), &users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)

	fmt.Println()
}
