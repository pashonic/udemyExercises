package main

import "fmt"

type person struct {
	first       string
	last        string
	favIcecream []string
}

func main() {
	people := []person{
		{
			first:       "nick",
			last:        "green",
			favIcecream: []string{"chochlate", "ass"},
		}, {
			first:       "todd",
			last:        "wells",
			favIcecream: []string{"white", "mint"},
		},
	}
	fmt.Println(people)
}
