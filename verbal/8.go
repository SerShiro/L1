package verbal

import "fmt"

func Task8() {
	// make используется для создания мап, каналов и слайсов
	m := make(map[int]int)
	c := make(chan int, 5)
	s := make([]int, 5)
	// new используется для выделения памяти под новое значение и возвращает указатель на эту память
	ptr := new(int)
	fmt.Println(m, c, s)
	fmt.Printf("%T\n", ptr)
}
