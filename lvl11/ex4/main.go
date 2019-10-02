package ex4

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// e := errors.New("more coffee needed")
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

// Resolve is a function that provide solution of exercise 4 ninja-level-11
func Resolve() {
	fmt.Println("Exercise 4")

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

	fmt.Println()
}
