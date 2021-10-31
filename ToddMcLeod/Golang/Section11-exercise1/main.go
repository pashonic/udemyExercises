package main

import "fmt"

func main() {
	x := [5]int{4, 5, 7, 8, 42}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
