package main

import "fmt"

type cake int

func main() {
	var x cake
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
