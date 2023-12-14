package tasks

import (
	"fmt"
)

func Task19() {
	str := "главрыба"
	fmt.Println("Исходная строка: ", str)
	fmt.Println("Перевернутая строка ", reverseString(str))
}
func reverseString(str string) string {
	// делаем из строки слайс рун
	runes := []rune(str)
	// идем с начала и с конца слайса, меняя местами
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// преобразуем обратно в строку
	str = string(runes)
	return str
}
