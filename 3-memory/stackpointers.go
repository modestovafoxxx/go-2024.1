package main

func main() {
	x := new(int)
	y := powX(*x) // тоже висит в стеке
	println(y)
}

func powX(x int) int {
	return x * x
}
