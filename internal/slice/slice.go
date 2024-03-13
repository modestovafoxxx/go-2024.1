package slice

// AppendElement добавляет элемент в конец слайса.
func AppendElement(ints []int, elem int) []int {
	// Код тут
	ints = append(ints, elem)
	return ints
}

// RemoveElement удаляет последний элемент из слайса. Если массив пуст, функция возвращает оригинальный пустой массив.
func RemoveElement(ints []int) []int {
	// Код тут
	if len(ints) == 0 {
		return ints
	}
	last_el := len(ints) - 1
	ints = ints[:last_el]
	return ints
}

// AddOneToAll увеличивает каждый элемент массива на единицу.
func AddOneToAll(ints []int) []int {
	// Код тут
	for i := 0; i <= len(ints)-1; i++ {
		ints[i]++
	}
	return ints
}
