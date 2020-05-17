package main

import "fmt"

func main() {
	a := 42 == 42
	b := 41 <= 42
	c := 41 >= 40
	d := 40 != 40
	e := 3 < 10
	f := 20 > 30

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
