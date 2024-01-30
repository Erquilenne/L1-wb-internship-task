package task6

/*
Реализовать все возможные способы остановки выполнения горутины.

Преимущества и недостатки различных способов остановки выполнения горутин:

1.Использование канала для отправки сигнала о завершении:

Преимущества: Прост в использовании, позволяет отправлять дополнительную информацию о причине остановки.
Недостатки: Возможно случайное закрытие закрытие канала из нескольких мест, что приведет к ошибкам.

2.спользование контекста для отмены выполнения:

Преимущества: Позволяет передавать информацию об отмене выполнения через иерархию вызовов.
Недостатки: Может потребовать изменения существующего кода для поддержки контекста.

3.Использование переменной-флага для контроля выполнения:

Преимущества: Прост в использовании и понимании, не требует дополнительных зависимостей.
Недостатки: Может привести к гонкам данных при одновременном доступе из нескольких горутин.

4.Использование функции runtime.Goexit() для немедленной остановки горутины:

Преимущества: Позволяет немедленно остановить выполнение горутины.
Недостатки: Может привести к непредсказуемому поведению и потере ресурсов, если не осуществляется корректная очистка.
*/

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Run() {
	// Использование канала для отправки сигнала о завершении
	stopChan := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("Received stop signal. Exiting goroutine.")
				return
			default:
				fmt.Println("Goroutine is running...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stopChan) // отправляем сигнал остановки

	// Использование контекста для отмены выполнения
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled. Exiting goroutine.")
				return
			default:
				fmt.Println("Goroutine is running...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel() // отменяем выполнение

	// Использование переменной-флага для контроля выполнения
	var stopFlag bool
	go func() {
		for !stopFlag {
			fmt.Println("Goroutine is running...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Stop flag is set. Exiting goroutine.")
	}()

	time.Sleep(3 * time.Second)
	stopFlag = true // устанавливаем флаг остановки

	// Использование функции runtime.Goexit() для немедленной остановки горутины
	go func() {
		defer fmt.Println("Exiting goroutine.")
		for {
			fmt.Println("Goroutine is running...")
			time.Sleep(1 * time.Second)
			if time.Now().After(time.Now().Add(3 * time.Second)) {
				runtime.Goexit() // немедленная остановка горутины
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
