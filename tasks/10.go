package tasks

import "fmt"

func Task10() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// создаем мапу, где ключ будет десяток температуры, а значение массивом из температур
	tempMap := map[int][]float32{}
	for _, v := range temp {
		tempMap[getDecade(v)*10] = append(tempMap[getDecade(v)*10], v)
	}
	fmt.Println(tempMap)
}

// получаем цифру десятка двузначного числа
func getDecade(num float32) int {
	return int(num) / 10
}
