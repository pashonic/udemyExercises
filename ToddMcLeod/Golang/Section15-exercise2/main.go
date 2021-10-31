package main

func main() {
	ii := []int{1, 2, 3, 4, 5}
	println(foo(ii...))
	println(bar(ii))
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
