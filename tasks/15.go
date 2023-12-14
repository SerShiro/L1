package tasks

import "fmt"

var justString string

func Task15() {
	someFunc()
	fmt.Println(justString)
}
func someFunc() {
	/*
		в изначальном куске проблема заключается в том, что мы хотим взять только небольшую часть строки,
		но когда используем слайс, justString ссылается на v, значит что мы будем хранить ссылку на большую
		строку, хотя нужно нам только 100 первых символов
		Изначальный кусок:
		v := createHugeString(1 << 10)
		justString = v[:100]
	*/
	v := createHugeString(1 << 10)
	// создаем новый слайс, и туда копируем первые 100 элементов строки, так получается что мы создали новый объект
	// и у него свое место в памяти
	justStringNew := make([]byte, 100)
	copy(justStringNew, v[:100])
	justString = string(justStringNew)
}

func createHugeString(size int) string {
	result := ""
	for i := 0; i < size; i++ {
		result += "HugeString"
	}
	return result
}
