package main

func main() {
	ch := make(chan int)

	go func(sc chan<- int) {
		for i := 0; i < 100; i++ {
			ch <- i

		}
		close(ch)
	}(ch)

	for v := range ch {
		println("Value:", v)
	}

	println("done")
}
