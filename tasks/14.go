package tasks

import "fmt"

func Task14() {
	// объявляем переменную типа interface, потом ей присваиваем значения int, string, bool, chan int
	var a interface{}
	a = 1
	// %T выводит тип переменной
	fmt.Printf("%T\n", a)
	a = "str"
	fmt.Printf("%T\n", a)
	a = true
	fmt.Printf("%T\n", a)
	a = make(chan int)
	fmt.Printf("%T\n", a)
}
