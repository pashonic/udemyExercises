package main

func main() {
	println(foo())
	x, y := bar()
	println(x)
	println(y)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "hello"
}
