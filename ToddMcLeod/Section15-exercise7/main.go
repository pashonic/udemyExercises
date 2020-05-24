package main

import "fmt"

func main() {
	square := func(num int) int {
		return num * num
	}

	v := square(5)

	fmt.Println(v)
}
