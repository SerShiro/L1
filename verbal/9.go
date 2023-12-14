package verbal

import "fmt"

func Task9() {
	// 3 способа задать слайс
	slice1 := []int{}
	var slice2 []int
	slice3 := make([]int, 0)
	fmt.Println(slice1, slice2, slice3)
	// 3 способа задания мапы, есть нюанс при задании map2, она является неинициализированной, при чтении она ведет себя
	// как пустая мапа, но при попытки записи будет паника, нужно инициализировать ее, например, через make
	map1 := map[int]int{}
	var map2 map[int]int
	map3 := make(map[int]int)
	fmt.Println(map1, map2, map3)
}
