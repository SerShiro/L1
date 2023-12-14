package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task5() {
	// Создаем канал для передачи данных и канал для завершения
	dataChan := make(chan int)
	doneChan := make(chan struct{})

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Запускаем горутину для чтения из канала, пока doneChan не закрыт
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case data := <-dataChan:
				fmt.Println("Прочитано из канала:", data)
			case <-doneChan:
				return
			}
		}
	}()

	// Запускаем горутину для отправки данных в канал каждые пол секунды, пока doneChan не закрыт
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(dataChan)
		for i := 1; ; i++ {
			time.Sleep(time.Millisecond * 500)
			select {
			case dataChan <- i:
			case <-doneChan:
				return
			}
		}
	}()

	// Ждем N секунд перед завершением программы
	N := 3
	fmt.Printf("Ожидание %d секунд...\n", N)
	time.Sleep(time.Second * time.Duration(N))
	// Закрываем канал и ждем завершения горутин
	close(doneChan)
	wg.Wait()

	// Завершаем программу
	fmt.Println("Программа завершена.")
}
