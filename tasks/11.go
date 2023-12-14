package tasks

import "fmt"

func Task11() {
	// создаем 2 мапы с ключами int и значениями bool(true что ключ есть), пересечения ищем по значениям ключей
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true, 7: true}
	fmt.Printf("Первое множество: ")
	printSet(set1)
	fmt.Printf("Второе множество: ")
	printSet(set2)
	fmt.Printf("Пересечение: ")
	printSet(intersection(set1, set2))

}

// проходимся по одной мапе и ищем есть ли в другой такие же значения, общие добавляем в ответ
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := map[int]bool{}
	for k := range set1 {
		if set2[k] {
			result[k] = true
		}
	}
	return result
}

// вывод значения ключей в строку
func printSet(set map[int]bool) {
	for k := range set {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")

}
