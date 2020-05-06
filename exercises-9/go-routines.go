package main

import (
	"fmt"
	"runtime"
	"sync"
)

// This code has a race condition
// During the loop, different go routines are trying to
// access the same variable counter
// and trying to write to it.
// This results in the variable never actually reaching
// the value of 100.

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Gouroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
