package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(cs chan<- int) {

	for i := 0; i < 100; i++ {
		cs <- i
	}
	close(cs)
}

func receive(cr <-chan int) {
	for v := range cr {
		fmt.Println("the value received from the channel:", v)
	}
}
