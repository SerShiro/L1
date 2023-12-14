package tasks

import (
	"fmt"
	"math/rand"
)

func Task13() {
	// случайно получаем задаем 2 числа
	num1 := rand.Intn(100)
	num2 := rand.Intn(100)

	fmt.Printf("До смены:\n\ta = %d, b = %d", num1, num2)
	// используем только исзначальные переменные и операторы "+" и "-"
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Printf("\nПосле смены:\n\ta = %d, b = %d\n", num1, num2)
}
