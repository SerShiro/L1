package tasks

import (
	"context"
	"fmt"
	"time"
)

func Task6() {
	go stopWithChan()
	go stopWithContext()
	go stopWithReturn()
	time.Sleep(time.Second * 4)
}

func stopWithChan() {
	// создаем канал
	ch := make(chan struct{})
	fmt.Println("Запускаем горутину с каналом!")
	// в горутине выводим сообщение о работе, пока канал открыт
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			select {
			case <-ch:
				fmt.Println("Горутина с каналом завершена!")
				return
			default:
				fmt.Println("Горутина с каналом работает...")
			}
		}
	}()
	time.Sleep(time.Second * 3)
	close(ch) // закрываем канал
	time.Sleep(time.Millisecond * 600)
}

func stopWithContext() {
	// создаем конекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Запускаем горутину с контекстом!")
	// аналогично вариацией с каналом, ждем пока пока конекст не завершится
	go func() {
		for {
			time.Sleep(time.Millisecond * 400)
			select {
			case <-ctx.Done():
				fmt.Println("Горутина с контекстом завершена!")
				return
			default:
				fmt.Println("Горутина с контекстом работает...")
			}
		}
	}()
	time.Sleep(time.Second * 3)
	cancel() // отменяем контекст
	time.Sleep(time.Millisecond * 600)
}

func stopWithReturn() {
	fmt.Println("Запускаем горутину с return!")
	// просто выводим данные пока не встречено условие выхода, выходим через return, так же вместо return можно использвать runtime.Goexit()
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Millisecond * 300)
			if i == 10 {
				fmt.Println("Горутина с return завершена!")
				return
			}
			fmt.Println("Горутина с return работает...")

		}
	}()
}
