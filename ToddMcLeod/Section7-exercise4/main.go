package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%v\t%b\t%X\n", a, a, a)
	a = a << 1
	fmt.Printf("%v,\t%b\t%X\n", a, a, a)
}
