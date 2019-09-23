package ex3

import (
	"fmt"
	"runtime"
	"sync"
)

// Resolve is a function that provide solution of exercise 3 ninja-level-9
func Resolve() {
	fmt.Println("Exercise 3")

	var wg sync.WaitGroup

	counter := 0
	iterations := 100
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			fmt.Println("Current value of counter is", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)

	fmt.Println()
}
