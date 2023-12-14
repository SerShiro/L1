package tasks

import (
	"fmt"
	"sync"
)

func Task7() {
	// создаем мапу и мьютекс
	myMap := make(map[string]int)
	var mu sync.Mutex

	// Количество горутин для конкурентной записи и WaitGroup
	numGoroutines := 5
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	// запускаем горутины, каждая из них создает по 10 элементов для мапы, используетя мьютекст для потокобезопасности
	for i := 0; i < numGoroutines; i++ {
		go func(num int) {
			defer wg.Done()

			// Конкурентная запись в мапу
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("Key%d", num*10+j)
				value := num*10 + j

				mu.Lock()
				myMap[key] = value
				mu.Unlock()
			}
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим результат
	fmt.Println("Размер мапы:", len(myMap))
	for k, v := range myMap {
		fmt.Println("[", k, "] = ", v)
	}
}
