package ex2

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error during marshalling %v in toJSON: %v", a, err)
	}
	return bs, nil
}

// Resolve is a function that provide solution of exercise 2 ninja-level-11
func Resolve() {
	fmt.Println("Exercise 2")

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

	fmt.Println()
}
