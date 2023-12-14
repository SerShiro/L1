package verbal

import (
	"fmt"
	"sync"
	"time"
)

// создаем стуктуру счетчик с Mutex
type counterMutex struct {
	sync.Mutex
	value int
}

// метод увеличиваем счетчик и блокирует доступ к значению, пока происходит запись
func (cm *counterMutex) increment() {
	cm.Lock()
	defer cm.Unlock()
	time.Sleep(time.Nanosecond)
	cm.value++
}

// метод возвращает значение счетчика и блокирует доступ к значению, пока происходит чтение
func (cm *counterMutex) getValue() int {
	cm.Lock()
	defer cm.Unlock()
	time.Sleep(time.Nanosecond)
	return cm.value
}

// создаем стуктуру счетчик с RWMutex
type counterRWMutex struct {
	sync.RWMutex
	value int
}

// метод увеличиваем счетчик и блокирует доступ к значению, пока происходит запись
func (crwm *counterRWMutex) increment() {
	crwm.Lock()
	defer crwm.Unlock()
	time.Sleep(time.Nanosecond)
	crwm.value++
}

// метод возвращает значение счетчика и значение блокируется только для чтение, т.е одновременно несколько горутин могут получить значение
func (cmrw *counterRWMutex) getValue() int {
	cmrw.RLock()
	defer cmrw.RUnlock()
	time.Sleep(time.Nanosecond)
	return cmrw.value
}
func Task3() {
	// основная разница Mutex и RWMutex в том, что в RWMutex мы можем сделать как блокировку на запись, так и на чтение, т.е
	// когда у нас несколько горутин только читают и никак не изменяют какое-то значение, RWMutex будет работать быстрее
	counterM()
	counterRWM()
}

func counterM() {
	start := time.Now()
	counter := counterMutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 50; i++ {

		go func() {
			defer wg.Done()

			counter.getValue()
		}()

		go func() {
			defer wg.Done()

			counter.increment()
		}()

	}
	wg.Wait()
	fmt.Println("Чтение и запись с Mutex для 50 горутин заняло: ", time.Since(start).Seconds())
}

func counterRWM() {
	start := time.Now()
	counter := counterRWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 50; i++ {

		go func() {
			defer wg.Done()

			counter.getValue()
		}()

		go func() {
			defer wg.Done()

			counter.increment()
		}()

	}
	wg.Wait()
	fmt.Println("Чтение и запись с RWMutex для 50 горутин заняло: ", time.Since(start).Seconds())
}
