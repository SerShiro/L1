package verbal

import "fmt"

func Task7() {
	// создаем мапу и добавляем элементы из примера
	myMap := map[int]int{}
	myMap[0] = 1
	myMap[1] = 124
	myMap[2] = 281
	// при итерации по мапе мы не можем точно сказать в каком порядке будут считываться элементы, потому что в основе
	// мап лежит хэш-таблицы, а в них порядок зависит от хэш-функций. Так что вывод в данном примере будет разный от запуска к запуску
	for k, v := range myMap {
		fmt.Printf("[%d] = %d ", k, v)
	}
	fmt.Println()
}
