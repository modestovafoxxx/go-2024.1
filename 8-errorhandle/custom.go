package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrCustom = fmt.Errorf("my custom error") // custom errors

func main() {
	customErr := ErrCustom

	var err error

	err = customErr // наша кастомная ошибка - как обычная ошибка golang

	_, err = os.ReadFile("not-exists.txt")

	err = fmt.Errorf("%v, %v", err, ErrCustom)

	if errors.As(err, &ErrCustom) {
		fmt.Println("there is wrapping!")
	}
}
