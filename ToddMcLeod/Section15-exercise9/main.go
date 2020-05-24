package main

import "fmt"

func main() {
	square := func(num int) int {
		return num * num
	}
	v := callme(square)
	fmt.Println(v)
}

func callme(fun func(int) int) int {
	v := fun(6)
	return v
}
