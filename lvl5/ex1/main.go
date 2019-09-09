package ex1

import (
	"fmt"
)

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

// Resolve is a function that provide solution of exercise 1 ninja-level-5
func Resolve() {
	fmt.Println("Exercise 1")

	people := []person{
		person{
			firstName: "Antonio",
			lastName:  "Margheriti",
			favFlavors: []string{
				"chocolate",
				"martini",
				"rum",
			},
		},
		person{
			firstName: "Enzo",
			lastName:  "Gorlomi",
			favFlavors: []string{
				"strawberry",
				"vanilla",
				"capuccino",
			},
		},
	}

	for _, p := range people {
		fmt.Println(p.firstName, p.lastName)
		for _, flavour := range p.favFlavors {
			fmt.Printf("%s ", flavour)
		}
		fmt.Println()
		fmt.Println()
	}

	fmt.Println()
}
