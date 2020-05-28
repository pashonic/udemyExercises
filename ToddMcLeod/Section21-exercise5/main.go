package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	runners := 100
	var wg sync.WaitGroup

	wg.Add(runners)
	for i := 0; i < runners; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			r := atomic.LoadInt64(&counter)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)
}
