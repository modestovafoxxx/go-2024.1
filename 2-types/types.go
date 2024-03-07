package main

import (
	"fmt"
	"image/color"
)

type Fruit struct {
	Color color.Color
}

type Apple struct {
	Fruit // вложенный тип
	Age   int
}

type Banana struct {
	Fruit
	Country string
}

func main() {
	a := Apple{ // объявляем новый объект типа Apple
		Fruit: Fruit{color.Black},
		Age:   3,
	}

	b := &Banana{ // объявляем новый объект типа Banana. b - будет указателем на этот объект
		Fruit:   Fruit{color.Black}, // вложенный тип - фрукт
		Country: "Africa",           // Банан африканский
	}

	fmt.Println(a, b)
}
