package verbal

import (
	"fmt"
	"sync"
)

func Task11() {
	wg := sync.WaitGroup{}
	// если оставить все без изменений, то программа выводит числа от 0 до 4, в случайном порядке и происходит дед лок
	// все из-за того, что внутрь горутин передается wg sync.WaitGroup, в каждой горутине свое копия и у каждой копии уменьшается счетчик на 1
	// группа в мейне остается без изменений
	// чтобы все работало корректно, можно передавать указатель
	for i := 0; i < 5; i++ {
		wg.Add(1)
		//go func(wg sync.WaitGroup, i int) {
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
			//}(wg, i)
		}(&wg, i)

	}
	wg.Wait()
	fmt.Println("exit")
}
