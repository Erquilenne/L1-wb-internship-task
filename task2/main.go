package task2

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

import (
	"fmt"
	"sync"
)

func Run() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для передачи результатов
	resultChannel := make(chan int)

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, resultChannel, &wg)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Читаем результаты из канала и выводим их
	for result := range resultChannel {
		fmt.Println(result)
	}
}

// Функция для вычисления квадрата числа и отправки результата в канал
func calculateSquare(number int, resultChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	resultChannel <- square
}
