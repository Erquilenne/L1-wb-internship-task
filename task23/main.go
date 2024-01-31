package task23

//Удалить i-ый элемент из слайса.

import "fmt"

func Run() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2 // Индекс элемента, который нужно удалить

	// Удаление i-го элемента из слайса
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice) // Результат: [1 2 4 5]
}
