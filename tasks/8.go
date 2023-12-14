package tasks

import "fmt"

func Task8() {
	number := int64(10) // Пример числа
	result := setBit(number, 0)

	fmt.Printf("Исходное значение: %d\n", number)
	fmt.Printf("Инвертируем 0ой бит: %d\n", result)

	number = int64(15) // Пример числа
	result = setBit(number, 3)

	fmt.Printf("Исходное значение: %d\n", number)
	fmt.Printf("Инвертируем 3ий бит: %d\n", result)
}

func setBit(value int64, position int) int64 {
	// делаем маску, сначала сдвигаем 1 на position потом побитовое отрицание, то есть все будет 1 кроме нашей позиции
	mask := ^(1 << position)
	// очищаем значение, побитовая конъюнкция, все значения кроме нашей позиции остатся
	clearedValue := value & int64(mask)
	// инвертируем значение которое получили
	var bitValue int
	if getBit(value, position) == 1 {
		bitValue = 0
	} else {
		bitValue = 1
	}

	// полученное значение ставим по позиции
	newValue := clearedValue | int64(bitValue<<position)
	return newValue
}

// получаем значение по нашей позиции 1 или 0
func getBit(value int64, position int) int {
	return int((value >> position) & 1)
}
