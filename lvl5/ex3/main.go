package ex3

import (
	"fmt"
)

type vehicle struct {
	doors byte
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

// Resolve is a function that provide solution of exercise 3 ninja-level-5
func Resolve() {
	fmt.Println("Exercise 3")

	truck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	sedan := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}
	fmt.Printf("%T\t%v\n", truck, truck)
	fmt.Printf("%T\t%v\n", sedan, sedan)
	fmt.Println(truck.doors)
	fmt.Println(sedan.doors)

	fmt.Println()
}
