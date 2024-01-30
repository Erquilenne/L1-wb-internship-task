package task9

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

import "fmt"

func Run() {
	// Создаем каналы для передачи данных
	input := make(chan int)
	output := make(chan int)

	// Горутина для записи чисел в первый канал
	go func() {
		defer close(input)
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			input <- num
		}
	}()

	// Горутина для умножения чисел и записи результата во второй канал
	go func() {
		defer close(output)
		for num := range input {
			output <- num * 2
		}
	}()

	// Вывод результата в stdout
	for result := range output {
		fmt.Println(result)
	}
}
