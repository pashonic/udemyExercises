package main

import "fmt"

func main() {
	m := map[string][]string{
		"James":           []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	for person, values := range m {
		fmt.Printf("Name: %v\n", person)
		for i, like := range values {
			fmt.Printf("\tFor index %v there is like %v\n", i, like)
		}
	}
}
