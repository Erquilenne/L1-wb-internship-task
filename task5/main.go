package task5

/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

import (
	"fmt"
	"time"
)

func sender(dataChannel chan<- int) {
	for i := 0; i < 10; i++ {
		dataChannel <- i
		time.Sleep(1 * time.Second) // Send a value every second
	}
	close(dataChannel)
}

func receiver(dataChannel <-chan int) {
	for {
		data, ok := <-dataChannel
		if !ok {
			break
		}
		fmt.Println("Received:", data)
	}
}

func Run() {
	dataChannel := make(chan int)
	go sender(dataChannel)
	go receiver(dataChannel)

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Time's up! Exiting...")
	}
}
