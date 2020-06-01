package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func(cs chan<- int) {
		cs <- 42
	}(ch)
	fmt.Println(<-ch)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", ch)
}
