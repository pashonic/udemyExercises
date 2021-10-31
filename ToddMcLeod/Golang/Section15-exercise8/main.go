package main

func main() {
	f := giveDoubleFunc()
	v := f(5)
	println(v)
}

func giveDoubleFunc() func(int) int {
	return func(x int) int {
		return x + x
	}
}
