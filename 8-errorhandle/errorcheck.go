package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("/not-exists.txt")
	// данный подход - стандартный в языке, как правило такое ставят достаточно часто чтобы
	// обработать ошибки и понять что делать дальше
	if err != nil {
		fmt.Printf("cant read such file with error: %s\n", err)

		os.Exit(2)
	}

	return
}
