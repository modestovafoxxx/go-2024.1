package main

import (
	"errors"
	"fmt"
)

// Аналогично Is(), функция As(err error, target interface{}) проверяет, соответствует ли какая-либо ошибка в цепочке обернутых ошибок целевой.
// Разница в том, что эта функция проверяет, имеет ли ошибка определенный тип, в отличие от Is(),
// которая проверяет, является ли ошибка -  конкретным объектом ошибки.

const badInput = "abc"

var ErrBadInput = errors.New("bad input") // так же можно объявлять через пакет errors новую ошибку

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf("validateInput: %w", ErrBadInput)
	}
	return nil
}

func main() {
	input := badInput

	err := validateInput(input)
	if errors.Is(err, ErrBadInput) {
		fmt.Println("bad input error")
	}
}
