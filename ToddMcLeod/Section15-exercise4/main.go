package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("may name is %v %v and my age is %v", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "tom",
		last:  "greene",
		age:   33,
	}
	p1.speak()
}
