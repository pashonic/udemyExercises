package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	inc := 0
	runners := 100

	var mu sync.Mutex

	wg.Add(runners)
	for i := 0; i < runners; i++ {
		go func() {
			mu.Lock()
			v := inc
			runtime.Gosched()
			v++
			inc = v
			mu.Unlock()
			fmt.Println(inc)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", inc)
}
