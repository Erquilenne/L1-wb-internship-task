package task18

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
 По завершению программа должна выводить итоговое значение счетчика.
*/

import (
	"fmt"
	"sync"
)

// Структура счетчика
type Counter struct {
	mu    sync.Mutex
	value int
}

// Метод инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func Run() {
	counter := Counter{}
	var wg sync.WaitGroup
	numRoutines := 100

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счетчика:", counter.value) // Output: Итоговое значение счетчика: 100
}
