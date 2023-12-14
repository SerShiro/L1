package tasks

import (
	"fmt"
	"sync"
)

func Task3() {
	// создаем слайс с заданными числами, канал и группу для горутин
	buff := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	ch := make(chan int, len(buff))

	wg.Add(len(buff)) // говорим что будет горутин столько же, сколько чисел в слайсек
	for _, v := range buff {
		// проходимся по всем элементам слайса, в горутинах возвращаем в канал квадраты чисел и уменьшаем счетчик группы на 1
		go func(num int) {
			defer wg.Done()
			ch <- num * num
		}(v)
	}
	// ждем пока счетчик группы станет 0, потом закрываем канал
	go func() {
		wg.Wait()
		close(ch)
	}()
	result := 0
	// выводим в stdout сумму всех чисел из канала
	for v := range ch {
		result += v
	}
	fmt.Println(result)
}
