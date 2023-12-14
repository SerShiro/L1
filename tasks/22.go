package tasks

import (
	"fmt"
	"math/big"
)

func Task22() {
	a := uint64(2<<55) + 123
	b := uint64(2<<33) + 321
	fmt.Printf("Число a = %d\nЧисло b = %d\n", a, b)
	fmt.Println("a * b = ", mult(a, b))
	if b == 0 {
		fmt.Println("Деление на 0 невозможно!")
	} else {

		fmt.Println("a / b = ", div(a, b))
	}
	fmt.Println("a + b = ", add(a, b))
	fmt.Println("a - b = ", sub(a, b))
}

func mult(num1, num2 uint64) *big.Int {
	a := new(big.Int).SetUint64(num1)
	b := new(big.Int).SetUint64(num2)
	result := new(big.Int).Mul(a, b)
	return result
}

func div(num1, num2 uint64) *big.Int {
	a := new(big.Int).SetUint64(num1)
	b := new(big.Int).SetUint64(num2)
	result := new(big.Int).Div(a, b)
	return result
}

func add(num1, num2 uint64) *big.Int {
	a := new(big.Int).SetUint64(num1)
	b := new(big.Int).SetUint64(num2)
	result := new(big.Int).Add(a, b)
	return result
}

func sub(num1, num2 uint64) *big.Int {
	a := new(big.Int).SetUint64(num1)
	b := new(big.Int).SetUint64(num2)
	result := new(big.Int).Sub(a, b)
	return result
}
