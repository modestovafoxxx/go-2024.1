package main

func main() {
	x := getValue()
	println(*x / 2)
}

func getValue() *int {
	res := 10
	return &res
}
