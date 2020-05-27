package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go go1()
	go go2()
	wg.Wait()
}

func go1() {
	fmt.Println("go1!")
	wg.Done()
}

func go2() {
	fmt.Println("go2!")
	wg.Done()
}
