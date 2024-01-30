package task8

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import "fmt"

func setBit(num int64, pos uint, bit uint) int64 {
	mask := int64(1 << pos)
	if bit == 1 {
		return num | mask // устанавливаем i-й бит в 1
	}
	return num &^ mask // устанавливаем i-й бит в 0
}

func Run() {
	var num int64 = 10 // пример исходного числа
	pos := uint(2)     // позиция бита
	bit := uint(1)     // значение бита (1 или 0)

	result := setBit(num, pos, bit)
	fmt.Printf("Результат: %d\n", result)
}
