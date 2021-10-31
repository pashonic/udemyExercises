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

	m := map[string]person{
		people[0].last: people[0],
		people[1].last: people[1],
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favIcecream {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}
}
