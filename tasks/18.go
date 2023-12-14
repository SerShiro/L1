package tasks

import (
	"fmt"
	"sync"
)

func Task18() {
	// создаем экземпляр структуры и WaitGroup
	myCounter := counter{}
	wg := sync.WaitGroup{}
	// в цикле делаем горутины, которые инкрементят счетчик
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			myCounter.increment()
		}()
	}
	// дожидаемся всех горутин и выводим значение счетчика на экран
	wg.Wait()
	fmt.Printf("Значение счетчика = %d\n", myCounter.getCount())
}

// в структуре используем мьютекс для потокобезопасности
type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
