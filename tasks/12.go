package tasks

import "fmt"

func Task12() {
	// записываем данную последовательность в слайс и создаем мапу для множества
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	mySet := map[string]bool{}

	fmt.Printf("Данная последовательность: ")
	// проходимся по слайсу и добавляем уникальные значения в мапу
	for _, v := range str {
		if !mySet[v] {
			mySet[v] = true
		}
		fmt.Printf("%s ", v)
	}

	fmt.Printf("\nПолученное множество: ")
	printMySet(mySet)

}

// вывод значения ключей в строку
func printMySet(set map[string]bool) {
	for k := range set {
		fmt.Printf("%s ", k)
	}
	fmt.Printf("\n")

}
