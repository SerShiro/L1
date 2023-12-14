package tasks

import (
	"fmt"
	"strings"
)

func Task26() {
	str := "abcdH"
	str1 := "abCdefAaf"
	str2 := "aabcd"
	fmt.Println(str, "-", checkUnique(str))
	fmt.Println(str1, "-", checkUnique(str1))
	fmt.Println(str2, "-", checkUnique(str2))
}

func checkUnique(str string) bool {
	// приводим строку в нижнему регистру и записываем в слайс рун
	arr := []rune(strings.ToLower(str))
	// создаем мапу для хранение количества элементов в слайсе
	runeMap := make(map[rune]int)
	// пробегаемся по слайсу и записываем в мапу сколько раз встретился каждый элемент
	for _, v := range arr {
		runeMap[v]++
	}
	// пробегаемся по мапе, если какое-то значение больше 1, значит в строке есть повторения
	for k, _ := range runeMap {
		if runeMap[k] > 1 {
			return false
		}
	}
	return true
}
