package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{
		name: "tom",
	}
	fmt.Println(p.name)

	changeMe(&p)

	fmt.Println(p.name)

}

func changeMe(p *person) {
	p.name = "nick"
}
