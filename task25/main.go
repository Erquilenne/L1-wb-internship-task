package task25

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

// Функция Sleep
func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func Run() {
	fmt.Println("Начало")
	Sleep(3) // Приостановка выполнения на 3 секунды
	fmt.Println("Прошло 3 секунды")
}
