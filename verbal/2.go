package verbal

import "fmt"

// интерфейс является неким контрактом,то есть представлять собой набор методов, которые должна реализовывать
// структура, которая реализует интерфейс
type Walker interface {
	walk()
}

type man struct{}

func (m *man) walk() {
	fmt.Println("I'm walking")
}

func Task2() {
	// так же можем сделать переменную типа интерфейс, такая переменная можем принимать значение любого типа
	// это можно использовать, когда мы не знаем какой тип будет исползоваться, например, в функции
	var anyValue interface{}
	anyValue = 12
	fmt.Println(anyValue)
	anyValue = 1.1
	fmt.Println(anyValue)
	anyValue = "Hello"
	fmt.Println(anyValue)
	anyValue = 'A'
	fmt.Println(anyValue)
}
