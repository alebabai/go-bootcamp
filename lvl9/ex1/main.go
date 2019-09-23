package ex1

import (
	"fmt"
	"runtime"
	"sync"
)

// Resolve is a function that provide solution of exercise 1 ninja-level-9
func Resolve() {
	fmt.Println("Exercise 1")

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing two")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

	fmt.Println()
}
