package ex2

import (
	"fmt"
)

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

// Resolve is a function that provide solution of exercise 2 ninja-level-5
func Resolve() {
	fmt.Println("Exercise 2")

	p1 := person{
		firstName: "Antonio",
		lastName:  "Margheriti",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum",
		},
	}

	p2 := person{
		firstName: "Enzo",
		lastName:  "Gorlomi",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	people := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range people {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, flavour := range v.favFlavors {
			fmt.Println(i, flavour)
		}
		fmt.Println("-------")
	}

	fmt.Println()
}
