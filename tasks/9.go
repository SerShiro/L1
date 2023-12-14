package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task9() {
	// создаем 2 канала, слайс с числами и WaitGroup
	inputChan := make(chan int)
	outputChan := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := sync.WaitGroup{}

	// в горутине записываем числа из слайса в inputChan, после этого закрываем его
	go func() {
		defer close(inputChan)
		for _, v := range nums {
			time.Sleep(time.Millisecond * 200)
			inputChan <- v
		}
	}()

	wg.Add(1)
	// в горутине читаем inputChan, значения оттуда умножаем на 2 и пишем в outputChan и  закрываем его
	go func() {
		defer wg.Done()
		for v := range inputChan {
			outputChan <- v * 2
		}
		close(outputChan)
	}()

	wg.Add(1)
	// в горутине читаем outputChan и выводим значения в stdout
	go func() {
		defer wg.Done()
		for v := range outputChan {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
