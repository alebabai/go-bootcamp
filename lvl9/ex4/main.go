package ex4

import (
	"fmt"
	"sync"
)

// Resolve is a function that provide solution of exercise 4 ninja-level-9
func Resolve() {
	fmt.Println("Exercise 4")

	counter := 0
	iterations := 100

	var wg sync.WaitGroup
	wg.Add(iterations)

	var mutex sync.Mutex

	for i := 0; i < iterations; i++ {
		go func() {
			mutex.Lock()
			temp := counter
			temp++
			counter = temp
			fmt.Println("Current value of counter is", counter)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)

	fmt.Println()
}
