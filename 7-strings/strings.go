package main

import (
	"fmt"
	"strings"
)

func main() {
	example := "hello world!"

	for _, r := range example {
		fmt.Printf("%t\n", r)
		// rune() - int32 потому что там так же могут быть символы дичь - например китайские
	}

	rplc := strings.ReplaceAll(example, "!", "?")

	fmt.Println(rplc)
}
