package tasks

import "fmt"

func Task17() {
	// создаем слайс и сортируем его, для бинарного поиска необхода отсортированная поселдовательность
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	target := 5
	if binSearch(arr, target) != -1 {
		fmt.Printf("Искомое число %d находится под индексом %d\n", target, binSearch(arr, target))
	} else {
		fmt.Printf("Искомое число %d не найдено\n", target)
	}
	target = 4
	if binSearch(arr, target) != -1 {
		fmt.Printf("Искомое число %d находится под индексом %d\n", target, binSearch(arr, target))
	} else {
		fmt.Printf("Искомое число %d не найдено\n", target)
	}
}

// пока low < high, проверяем средний элемент равен искомому, если больше то ищем в правой части, иначе в левой
// если не найден возвращаем -1
func binSearch(arr []int, searchNum int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if searchNum == arr[mid] {
			return mid
		} else if searchNum < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
