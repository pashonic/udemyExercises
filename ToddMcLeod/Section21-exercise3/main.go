package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	inc := 0
	runners := 100

	wg.Add(runners)
	for i := 0; i < runners; i++ {
		go func() {
			v := inc
			runtime.Gosched()
			v++
			inc = v
			fmt.Println(inc)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}
