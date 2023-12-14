package tasks

import "fmt"

func Task23() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Исходный слайс :", arr)
	index := 2
	fmt.Println("После удаления элемента с индексом ", index, ":", removeOnIndex(arr, index))
}

// возвращаем новый слайс, в котором  записываем все элементы до индекса и после
func removeOnIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
