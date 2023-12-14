package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Task4() {
	const numWorkers = 5

	// Создаем канал для передачи данных
	dataChan := make(chan int, 10)

	// Создаем канал для сигнала завершения
	doneChan := make(chan struct{})

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	// Запускаем горутину для записи данных в канал, горутина завершится, когда закроется канал doneChan
	go func() {
		defer close(dataChan)

		for i := 0; ; i++ {
			select {
			case dataChan <- i: // Записываем данные в канал
			case <-doneChan: // Завершаем запись данных при получении сигнала завершения
				return
			}
		}
	}()
	// канал для сигнала прерывания
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)

	// если читаем из канал, то закрываем канал doneChan
	go func() {
		<-sigChan
		fmt.Println("\nПринят сигнал прерывания. Завершение работы...")
		close(doneChan)
	}()

	// Ожидаем завершения всех воркеров
	wg.Wait()
	fmt.Println("Главная программа завершена.")
}

func worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	//Каждый воркер раз в пол секунды получает данные из канала, если канал закрыт, то воркер завершается
	for {
		time.Sleep(time.Millisecond * 500)
		select {
		case data, ok := <-dataChan:
			if !ok {
				// Канал закрыт, завершаем работу воркера
				fmt.Printf("Воркер %d завершен.\n", id)
				return
			}
			fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		}
	}
}
