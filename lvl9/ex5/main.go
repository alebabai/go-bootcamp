package ex5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Resolve is a function that provide solution of exercise 5 ninja-level-9
func Resolve() {
	fmt.Println("Exercise 5")

	var counter int64
	iterations := 100

	var wg sync.WaitGroup
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			temp := atomic.LoadInt64(&counter)
			temp++
			atomic.StoreInt64(&counter, temp)
			fmt.Println("Current value of counter is", temp)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)

	fmt.Println()
}
