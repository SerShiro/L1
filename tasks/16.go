package tasks

import (
	"fmt"
)

func Task16() {
	// создаем два слайса
	arr1 := []int{5, 16, 13, 8, 6, 1, 2}
	arr2 := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Первый массив до сортировки: ", arr1)
	quickSort(arr1, 0, len(arr1)-1)
	fmt.Println("Первый массив после сортировки: ", arr1)
	fmt.Println("Второй массив до сортировки: ", arr2)
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Println("Второй массив после сортировки: ", arr2)
}

func quickSort(arr []int, low, high int) {
	// производим сортировку когда low < high
	if low < high {
		// получаем индекс опорного элемента
		pivotIndex := partition(arr, low, high)
		// вызываем функцию сортировки для левой и правой части от опорного элемента
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// берем последний элемент как опорный и двигаем Wall, слева Wall элементы меньше опорного, справа больше опорного
// возвращаем индекс первого справа элемента от Wall
func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
