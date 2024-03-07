package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()
	// можно использовать чтобы не забыть,
	// тогда по стеку когда кончается функция в которой объявлен defer будет - он выполнен
}
