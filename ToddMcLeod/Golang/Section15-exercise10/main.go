package main

import "fmt"

func main() {
	a := foo("a")
	b := foo("b")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(a())
}

func foo(n string) func() (string, int) {
	x := 0
	return func() (string, int) {
		x++
		return n, x
	}
}
