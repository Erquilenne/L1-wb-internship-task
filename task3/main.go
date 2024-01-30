package task3

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

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

	// Суммируем результаты из канала
	sum := 0
	for result := range resultChannel {
		sum += result
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}

// Функция для вычисления квадрата числа и отправки результата в канал
func calculateSquare(number int, resultChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	resultChannel <- square
}
