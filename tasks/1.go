package tasks

import "fmt"

// Human
// создание структуры, которая содержит 2 поля типа string и 2 метода: первый пустой, второй требует строку
type Human struct {
	name, secondName string
}

func (h *Human) walk() {
	fmt.Println("I'm walking")
}

func (h *Human) sayHello(name string) {
	fmt.Println("Hello", name)
}

// Action
// определяется структура Action с встраиванием Human
type Action struct {
	Human
}

func Task1() {
	a1 := Action{} // создаем экземпляр структуры Action
	// вызываем методы Human
	a1.walk()
	a1.sayHello("Ivan")
}
