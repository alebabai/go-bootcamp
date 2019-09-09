package ex4

import (
	"fmt"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-5
func Resolve() {
	fmt.Println("Exercise 4")

	song := struct {
		title  string
		verses []struct {
			quotes []struct {
				phrase string
			}
		}
	}{
		title: "Всё идёт по плану",
	}
	fmt.Println(song)
	fmt.Println(song.title)
	fmt.Println(song.verses)

	fmt.Println()
}
