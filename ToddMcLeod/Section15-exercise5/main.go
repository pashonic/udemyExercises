package main

import (
	"fmt"
	"math"
)

type square struct {
	L float64
}

type circle struct {
	R float64
}

func (s square) area() float64 {
	return s.L * s.L
}

func (s circle) area() float64 {
	return math.Pi * s.R * s.R
}

type shape interface {
	area() float64
}

func info(s shape) {
	area := s.area()
	fmt.Println(area)
}

func main() {
	circ := circle{
		R: 12.345,
	}

	squa := square{
		L: 15,
	}

	info(circ)
	info(squa)
}
