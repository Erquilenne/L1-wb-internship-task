package task7

//  Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

func Run() {
	data := make(map[string]int)
	var mutex sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", index)
			mutex.Lock()
			data[key] = index
			mutex.Unlock()
			fmt.Printf("Goroutine %d wrote to map\n", index)
		}(i)
	}
	wg.Wait()

	fmt.Println("Final map contents:", data)
}
