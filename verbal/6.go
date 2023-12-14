package verbal

import "fmt"

type calculation struct{}

// cоздаем метод структуры, который принимает 2 числа и выводит их сумму
func (c calculation) add(x, y int) {
	fmt.Println("Result:", x+y)
}

/* пробуем сделать перегрузку метода, хотим чтобы теперь 3 числа складывалось, но IDE не дает нам этого сделать
func (c calculation) add(x, y, z int) {
	fmt.Println("Result:", x+y+z)
}
*/

func Task6() {
	fmt.Println("Язык Go не поддерживает перегрузку")
}
