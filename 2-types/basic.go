package main

import (
	"fmt"
)

func main() {
	var i int = 3
	var u uint = 5
	var c string = "hello world!"

	fmt.Println(i, u, c)

	m := make(map[string]int)
	m["apple"] = 3

	_, exists := m["banana"]
	if !exists {
		fmt.Println("banana not exists")
	}

	appleCount, exists := m["apple"]
	if !exists {
		fmt.Println("apple not exists")
	} else {
		fmt.Println("apple exists")
		fmt.Println(appleCount)
	}
}
