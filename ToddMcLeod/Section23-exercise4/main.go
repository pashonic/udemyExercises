package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chcan int)

	go func () {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(cr <-chan int, q <-chan int) {
	for {
		select {
		case v := <-cr:
			fmt.Println("from the loop channel:", v)
		case <-q:
			fmt.Println("Quit message recieved")
			return
		}
	}
}