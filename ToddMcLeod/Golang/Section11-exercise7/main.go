package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	for i, _ := range x {
		fmt.Println("kick")
		for j, v := range x[i] {
			fmt.Println(j, v)
		}
	}
}
