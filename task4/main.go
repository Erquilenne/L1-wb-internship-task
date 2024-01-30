package task4

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChannel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Worker %d received close signal, exiting\n", id)
				return
			}
			fmt.Printf("Worker %d received data: %d\n", id, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancel signal, exiting\n", id)
			return
		}
	}
}

func Run(numWorkers int) {
	dataChannel := make(chan int)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start N worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, ctx, &wg)
	}

	// Write data to the channel (in the main goroutine)
	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i
		}
		close(dataChannel)
	}()

	// Wait for Ctrl+C signal to cancel the context
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	cancel() // Cancel the context to signal all workers to exit

	// Wait for all worker goroutines to finish
	wg.Wait()
}
