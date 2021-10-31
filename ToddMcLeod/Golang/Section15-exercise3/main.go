package main

func main() {
	defer foo()
	println("first")
}

func foo() {
	defer func() { println("very last") }()
	println("last")
}
