package main

import dog
import "fmt"

type canine struct {
	name string
	age int
}

func main() {
	fido := canine {
		name : "Fido",
		age : dog.Years(10)
	}
	fmt.Print
}